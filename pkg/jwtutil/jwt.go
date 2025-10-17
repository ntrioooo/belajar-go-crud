package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
	NewAccessToken(userID uint) (string, error)
	Parse(tokenStr string) (*jwt.Token, jwt.MapClaims, error)
}

type manager struct{ secret []byte }

func New(secret string) TokenManager { return &manager{secret: []byte(secret)} }

func (m *manager) NewAccessToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour * 30).Unix(),
	})
	return token.SignedString(m.secret)
}

func (m *manager) Parse(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) { return m.secret, nil })
	if err != nil {
		return nil, nil, err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return t, claims, nil
	}
	return t, nil, err
}
