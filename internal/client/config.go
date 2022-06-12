package client

import (
	"context"
	"fmt"

	"github.com/aveplen-bach/config-gateway/internal/model"
	"github.com/aveplen-bach/config-gateway/protos/config"
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
	if _, err := cc.cc.UpdateFacerecConfig(context.Background(), &config.FacerecConfig{
		Threshold: facerecConfig.Threshold,
	}); err != nil {
		return fmt.Errorf("could call config client rps: %w", err)
	}
	return nil
}
