package middleware

import (
	"fmt"
	"net/http"
	"github.com/ahdaan98/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthMiddleware(c *gin.Context){

	cfg,_:=config.LoadEnvVariables()

	tokenString, err := c.Cookie("client")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}

	// Decode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cfg.ACCESS_KEY_USER), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token, please log in"})
		c.Abort()
		return
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token, please log in"})
		c.Abort()
		return
	}

	fmt.Println("claims", claims)

	role, ok := claims["role"].(string)
	if !ok || role != "client" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "error in retrieving id"})
		c.Abort()
		return
	}

	c.Set("role", role)
	c.Set("id", int(id))

	c.Next()
}