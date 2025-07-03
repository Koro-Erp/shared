package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HasPermission(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawPerms, exists := c.Get("permissions")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "Permissions not found"})
			c.Abort()
			return
		}

		permList, ok := rawPerms.([]string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "Invalid permission format"})
			c.Abort()
			return
		}

		for _, p := range permList {
			if p == required {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "Permission denied"})
		c.Abort()
	}
}
