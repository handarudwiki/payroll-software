package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/handarudwiki/payroll-sistem/internal/models"
)

func GenerateToken(userId int, role models.UserRole, jwtSecret string) (string, error) {

	now := time.Now()

	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     now.Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(([]byte(jwtSecret)))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string, jwtSecret string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, err
}
