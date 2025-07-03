package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireScope(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		mapClaims := claims.(jwt.MapClaims)

		if mapClaims["type"] != "service" {
			c.JSON(403, gin.H{"code": 403, "message": "Only service tokens allowed"})
			c.Abort()
			return
		}

		scopeRaw := mapClaims["scope"]
		if scopeRaw == nil {
			c.JSON(403, gin.H{"code": 403, "message": "No scopes"})
			c.Abort()
			return
		}

		scopes := scopeRaw.([]interface{})
		for _, s := range scopes {
			if s == required {
				c.Next()
				return
			}
		}

		c.JSON(403, gin.H{"code": 403, "message": "Missing required scope"})
		c.Abort()
	}
}
