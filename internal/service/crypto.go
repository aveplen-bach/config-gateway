package service

import (
	"fmt"

	"github.com/aveplen-bach/config-gateway/internal/client"
)

type CryptoService struct {
	ac *client.AuthClient
}

func NewCryptoService(ac *client.AuthClient) *CryptoService {
	return &CryptoService{
		ac: ac,
	}
}

func (cs *CryptoService) Encrypt(userID uint, opentext []byte) ([]byte, error) {
	ciphertext, err := cs.ac.Encrypt(userID, opentext)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt due to client error: %w", err)
	}

	return ciphertext, nil
}

func (cs *CryptoService) Decrypt(userID uint, ciphertext []byte) ([]byte, error) {
	opentext, err := cs.ac.Decrypt(userID, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt due to client error: %w", err)
	}

	return opentext, nil
}
