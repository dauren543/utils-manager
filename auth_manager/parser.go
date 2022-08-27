package auth_manager

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"
)

func (m *authManager) ParseDefaultClaims(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m *authManager) NewRefreshToken() (string, error) {
	var (
		b = make([]byte, 32)
		r = rand.New(rand.NewSource(time.Now().Unix()))
	)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *authManager) CheckAuthToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(m.SigningKey),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		},
	})
}
