package auth_manager

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type authManager struct {
	SigningKey string
}

func NewAuthManager(signingKey string) *authManager {
	return &authManager{SigningKey: signingKey}
}

func (m *authManager) NewJwtWithDefaultClaims(ttl time.Duration, claims string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"sub": claims,
	})

	return token.SignedString([]byte(m.SigningKey))
}