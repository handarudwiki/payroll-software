package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
)

func AuthMiddleware(cfg config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		bearer := strings.Split(token, "Bearer ")

		if len(bearer) != 2 {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		token = bearer[1]

		claims, err := utils.ValidateToken(token, cfg.Secret)
		if err != nil {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("id", claims.UserId)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AuthorizationMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		isAuthorized := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			utils.ResponseError(c, "forbidden", http.StatusForbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}
