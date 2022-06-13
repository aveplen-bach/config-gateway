package service

import (
	"fmt"

	"github.com/aveplen-bach/config-gateway/internal/client"
	"github.com/aveplen-bach/config-gateway/internal/model"
	"github.com/sirupsen/logrus"
)

type ConfigService struct {
	cc *client.ConfigClient
}

func NewConfigService(cc *client.ConfigClient) *ConfigService {
	return &ConfigService{
		cc: cc,
	}
}

func (cs *ConfigService) UpdateFacerecConfig(frCfg model.FacerecConfig) error {
	logrus.Info("updating facerec config")
	if err := cs.cc.UpdateFacerecConfig(frCfg); err != nil {
		logrus.Warnf("could not update facerec config due to client error: %w", err)
		return fmt.Errorf("could not update facerec config due to client error: %w", err)
	}
	return nil
}
