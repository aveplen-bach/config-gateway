package client

import (
	"context"
	"fmt"

	"github.com/aveplen-bach/config-gateway/protos/auth"
	"github.com/sirupsen/logrus"
)

type AuthClient struct {
	ac auth.AuthenticationClient
}

func NewAuthClient(ac auth.AuthenticationClient) *AuthClient {
	return &AuthClient{
		ac: ac,
	}
}

func (ac *AuthClient) Encrypt(userID uint, opentext []byte) ([]byte, error) {
	logrus.Info("calling auth encrypt rpc")
	ciphertext, err := ac.ac.Encrypt(context.Background(), &auth.Opentext{
		Id:       uint64(userID),
		Contents: opentext,
	})

	if err != nil {
		logrus.Warn("could not call auth encrypt rpc: %w", err)
		return nil, fmt.Errorf("could not call auth encrypt rpc: %w", err)
	}

	return ciphertext.Contents, nil
}

func (ac *AuthClient) Decrypt(userID uint, ciphertext []byte) ([]byte, error) {
	logrus.Info("calling auth decrypt rpc")
	opentext, err := ac.ac.Decrypt(context.Background(), &auth.Ciphertext{
		Id:       uint64(userID),
		Contents: ciphertext,
	})

	if err != nil {
		logrus.Warn("could not call auth decrypt rpc: %w", err)
		return nil, fmt.Errorf("could not call auth decrypt rpc: %w", err)
	}

	return opentext.Contents, nil
}

func (ac *AuthClient) GetNextSynPackage(userID uint, syn []byte) ([]byte, error) {
	logrus.Info("calling auth get next syn package rpc")
	res, err := ac.ac.GetNextSynPackage(context.Background(), &auth.SynPackage{
		Id:       uint64(userID),
		Contents: syn,
	})

	if err != nil {
		logrus.Warnf("could not call auth get next syn package rpc: %w", err)
		return nil, fmt.Errorf("could not call auth get next syn package rpc: %w", err)
	}

	return res.Contents, nil
}
