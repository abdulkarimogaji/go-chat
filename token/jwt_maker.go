package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	return &JwtMaker{}, nil
}

func (m *JwtMaker) CreateToken(userId int, duration time.Duration) (string, error) {
	payload, err := NewPayload(userId, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(m.secretKey))
}
func (m *JwtMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInValidToken
		}
		return []byte(m.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, jwt.StandardClaims{}, keyFunc)
	if err != nil {
		if v_err, ok := err.(*jwt.ValidationError); ok && errors.Is(v_err.Inner, ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInValidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrInValidToken
	}
	return payload, nil
}
