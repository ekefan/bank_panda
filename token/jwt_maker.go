package token

///Lecture 21: Create and Verfiy JWT & PASETO token in Golang

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const minSecretKeySize = 32 //characters

// JWTMaker is a JSON web Token maker
type JWTMaker struct{
	secretKey string

}


func NewJWTMaker(secretkey string) (Maker, error) {
	if len(secretkey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %v characters", minSecretKeySize)
	}

	return &JWTMaker{secretkey}, nil
}


// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker)CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	return jwtToken.SignedString([]byte(maker.secretKey))

}

// VerifyToken checks if the token is valid or not, 
// Returns the Payload data on true and error on false
func (maker *JWTMaker)VerifyToken(token string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)
	if err  != nil {
		// Note  is v5 validation error did not exist
		vErr, ok := err.(*jwt.ValidationError) // to get whether token is Expired or not valid
		if ok && errors.Is(vErr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil,  ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
