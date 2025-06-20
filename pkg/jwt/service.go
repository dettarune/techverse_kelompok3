package jwt

import (
	"context"
	"fmt"
	"time"

	jwtLib "github.com/golang-jwt/jwt"
)

type Service struct {
	secretKey string
	issuer    string
}

func NewService(secretKey string) *Service {
	return &Service{
		secretKey: secretKey,
		issuer:    "todo-go",
	}
}

func (s *Service) ParseToken(ctx context.Context, tokenString string) (token *jwtLib.Token, err error) {
	return jwtLib.Parse(tokenString, func(token *jwtLib.Token) (any, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

func (s *Service) GenerateToken(ctx context.Context, data map[string]any, exp int64) (string, error) {
	claims := jwtLib.MapClaims(data)
	claims["exp"] = exp
	claims["iss"] = s.issuer
	claims["iat"] = time.Now().Unix()

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
