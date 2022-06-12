package client

import (
	"context"
	"fmt"

	"github.com/aveplen-bach/config-gateway/protos/auth"
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
	ciphertext, err := ac.ac.Encrypt(context.Background(), &auth.Opentext{
		Id:       uint64(userID),
		Contents: opentext,
	})

	if err != nil {
		return nil, fmt.Errorf("could not call auth client rpc method: %w", err)
	}

	return ciphertext.Contents, nil
}

func (ac *AuthClient) Decrypt(userID uint, ciphertext []byte) ([]byte, error) {
	opentext, err := ac.ac.Decrypt(context.Background(), &auth.Ciphertext{
		Id:       uint64(userID),
		Contents: ciphertext,
	})

	if err != nil {
		return nil, fmt.Errorf("could not call auth client rpc method: %w", err)
	}

	return opentext.Contents, nil
}

func (ac *AuthClient) GetNextSynPackage(userID uint, syn []byte) ([]byte, error) {
	res, err := ac.ac.GetNextSynPackage(context.Background(), &auth.SynPackage{
		Id:       uint64(userID),
		Contents: syn,
	})

	if err != nil {
		return nil, fmt.Errorf("could not call auth client rpc method: %w", err)
	}

	return res.Contents, nil
}
