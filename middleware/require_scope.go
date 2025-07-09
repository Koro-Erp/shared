package middleware

import (
	"net/http"

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

func RequireUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("user_claims"); !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "user token required"})
			return
		}
		c.Next()
	}
}


func RequireServiceToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("service_claims"); !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "service token required"})
			return
		}
		c.Next()
	}
}


func RequireBothTokens() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, userOK := c.Get("user_claims")
		_, serviceOK := c.Get("service_claims")
		if !userOK || !serviceOK {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "both service and user tokens required"})
			return
		}
		c.Next()
	}
}
