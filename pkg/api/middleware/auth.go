package middleware

import (
	"github.com/ahdaan98/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {

	cfg, _ := config.LoadEnvVariables()

	tokenString, err := c.Cookie("admin")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token, please log in"})
		c.Abort()
		return
	}

	// Decode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.ACCESS_KEY_ADMIN), nil
	})

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token, please log in"})
		c.Abort()
		return

	}

	c.Next()
}
