package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(publicKey interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var serviceOK, userOK bool
		requireUserToken := true // Default: user token is required

		// Parse service token
		authHeader := c.GetHeader("Authorization")
		// log.Println(authHeader)
		if strings.HasPrefix(authHeader, "Bearer ") {
			serviceTokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			serviceClaims := jwt.MapClaims{}
			if _, err := jwt.ParseWithClaims(serviceTokenStr, serviceClaims, func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			}); err == nil && serviceClaims["token_type"] == "service" {
				c.Set("service_claims", serviceClaims)
				serviceOK = true

				// Try to extract require_user_token as a bool
				if val, ok := serviceClaims["require_user_token"]; ok {
					
					if b, ok := val.(bool); ok {
						requireUserToken = b
					}
				}
			}
		}

		// Parse optional user token
		userHeader := c.GetHeader("X-User-Token")
		// log.Println(userHeader)
		if strings.HasPrefix(userHeader, "Bearer ") {
			userTokenStr := strings.TrimPrefix(userHeader, "Bearer ")
			userClaims := jwt.MapClaims{}
			if _, err := jwt.ParseWithClaims(userTokenStr, userClaims, func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			}); err == nil && userClaims["token_type"] == "user" {
				c.Set("user_claims", userClaims)
				userOK = true
			}
		}

		// Require at least one valid token
		if !serviceOK && !userOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code":400,"message": "valid service or user token required"})
			return
		}

		// If service requires user token, but it's missing
		if serviceOK && requireUserToken && !userOK {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code":400,"message": "user token required for this service"})
			return
		}

		c.Next()
	}
}
