package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/gocastsian/writino/errors"
)

func GenerateTokenPair(secret []byte, userID string) (map[string]string, error) {

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Subject:   userID,
	}
	refClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refClaims)

	encoded, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	refEncoded, err := refToken.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  encoded,
		"refresh_token": refEncoded,
	}, nil
}

func ParseToken(secret []byte, token string) (string, error) {

	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsed.Valid {
		return "", errors.ErrInvalidToken
	}

	claims := parsed.Claims.(*jwt.StandardClaims)
	return claims.Subject, nil
}
