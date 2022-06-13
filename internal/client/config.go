package client

import (
	"context"
	"fmt"

	"github.com/aveplen-bach/config-gateway/internal/model"
	"github.com/aveplen-bach/config-gateway/protos/config"
	"github.com/sirupsen/logrus"
)

type ConfigClient struct {
	cc config.ConfigClient
}

func NewConfigClient(cc config.ConfigClient) *ConfigClient {
	return &ConfigClient{
		cc: cc,
	}
}

func (cc *ConfigClient) UpdateFacerecConfig(facerecConfig model.FacerecConfig) error {
	logrus.Info("calling config update facerec config rpc")
	if _, err := cc.cc.UpdateFacerecConfig(context.Background(), &config.FacerecConfig{
		Threshold: facerecConfig.Threshold,
	}); err != nil {
		logrus.Warnf("could not call config update facerec config rpc: %w", err)
		return fmt.Errorf("could not call config update facerec config rpc: %w", err)
	}
	return nil
}
