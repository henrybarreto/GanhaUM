package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionAllowOnly(addresses ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var found bool
		for _, address := range addresses {
			if address == c.ClientIP() {
				found = true

				break
			}
		}

		if !found {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func PermissionBlockOnly(addresses ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, address := range addresses {
			if address == c.ClientIP() {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		c.Next()
	}
}

func PermissionBlockAll(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}
