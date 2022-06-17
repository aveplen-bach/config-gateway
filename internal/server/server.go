package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aveplen-bach/config-gateway/internal/client"
	"github.com/aveplen-bach/config-gateway/internal/config"
	"github.com/aveplen-bach/config-gateway/internal/handler"
	"github.com/aveplen-bach/config-gateway/internal/middleware"
	"github.com/aveplen-bach/config-gateway/internal/service"
	"github.com/aveplen-bach/config-gateway/protos/auth"
	confpb "github.com/aveplen-bach/config-gateway/protos/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(cfg config.Config) {
	// ============================= auth client ==============================
	logrus.Info("connecting to auth service")
	authTimeout, authCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer authCancel()

	authCC, err := grpc.DialContext(authTimeout, cfg.AuthClient.Addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logrus.Warn(fmt.Errorf("failed to connecto to %s: %w", cfg.AuthClient.Addr, err))
	}

	authClient := client.NewAuthClient(auth.NewAuthenticationClient(authCC))

	// ============================= conf client ==============================
	logrus.Info("connecting to config service")
	confTimeout, confCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer confCancel()

	confCC, err := grpc.DialContext(confTimeout, cfg.ConfigClient.Addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logrus.Warn(fmt.Errorf("failed to connecto to %s: %w", cfg.ConfigClient.Addr, err))
	}

	configClient := client.NewConfigClient(confpb.NewConfigClient(confCC))

	// ================================ service ===============================
	tokentService := service.NewTokenService(cfg, authClient)
	configService := service.NewConfigService(configClient)
	cryptoService := service.NewCryptoService(authClient)

	// ================================ router ================================
	logrus.Info("creating router")
	r := gin.Default()
	r.Use(middleware.Cors())

	encr := r.Group("/api/config/encr")
	encr.Use(middleware.Incremental(tokentService))
	encr.Use(middleware.Encryption(tokentService, cryptoService))

	// ================================ routes ================================
	logrus.Info("registering routes")
	encr.POST("/facerec", handler.UpdateFacerecConfig(configService))

	r.GET("/cg/health/live", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// =============================== shutdown ===============================
	srv := &http.Server{
		Addr:    cfg.ServerConfig.Addr,
		Handler: r,
	}

	go func() {
		logrus.Infof("listening: %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Warn(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Warn("shutting down server...")

	ctx, authCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer authCancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown:", err)
	}

	logrus.Warn("server exited")
}
