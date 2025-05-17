package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/handarudwiki/payroll-sistem/internal/models"
)

type JWTClaims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, role models.UserRole, jwtSecret string) (string, error) {

	now := time.Now()

	claims := JWTClaims{
		UserId: userId,
		Role:   string(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string, jwtSecret string) (claims JWTClaims, err error) {
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return claims, err
	}

	return claims, err
}
