package handler

import (
	"net/http"

	"github.com/aveplen-bach/config-gateway/internal/model"
	"github.com/aveplen-bach/config-gateway/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpdateFacerecConfig(cs *service.ConfigService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCfg model.FacerecConfig
		if err := c.BindJSON(&newCfg); err != nil {
			logrus.Warnf("could not bind json: %w", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := cs.UpdateFacerecConfig(newCfg); err != nil {
			logrus.Warnf("could not update facerec config: %w", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}

		logrus.Info("facerec config updated")
		c.JSON(http.StatusOK, gin.H{
			"info": "facerec config updated",
		})
	}
}
