package handler

import (
	"net/http"

	"github.com/aveplen-bach/config-gateway/internal/model"
	"github.com/aveplen-bach/config-gateway/internal/service"
	"github.com/gin-gonic/gin"
)

func UpdateFacerecConfig(cs *service.ConfigService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCfg model.FacerecConfig
		if err := c.BindJSON(&newCfg); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := cs.UpdateFacerecConfig(newCfg); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"info": "facerec config updated",
		})
	}
}
