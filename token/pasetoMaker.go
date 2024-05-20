package token

import (
	
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"fmt"
	"time"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto *paseto.V2

	// since it is being used locally for backend api
	// we would use symmetric encrytpion to encrypt the token payload
	symmetricKey []byte
}


func NewPasetoMaker(symmetrickey string) (Maker, error) {
	if len(symmetrickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key seiz: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symmetricKey: []byte(symmetrickey),
	}

	return maker, nil
}


//Create Token creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

//
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil

}