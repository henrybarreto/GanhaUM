package middleware

import (
	"fmt"
	"ganhaum.henrybarreto.dev/internal/api/http/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getTokenFromHeader(header string) (string, error) {
	if len(header) == 0 {
		return "", fmt.Errorf("no authorization header provided")
	}

	splitted := strings.Split(header, " ")
	if len(splitted) != 2 || splitted[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header provided")
	}

	return splitted[1], nil
}

func AuthenticationJWT(c *gin.Context) {
	authorization, err := getTokenFromHeader(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if ok, err := jwt.CheckJWT(authorization); err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
