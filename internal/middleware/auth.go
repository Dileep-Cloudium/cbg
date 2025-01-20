package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/portal_api"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// List of valid API keys
var validAPIKeys = []string{
	"ng-api-key-f7d9a2e1b3c5",
	"ng-api-key-8h4k9m2p5n7",
	"ng-api-key-3x6y9w2q4v8",
	"ng-api-key-5t8u1j4r7d9",
	"ng-api-key-2b5n8m1h4k7",
}

// AuthMiddleware checks for a valid token in the request headers
func AuthMiddleware() gin.HandlerFunc {
	dbClient, _ := db.NewDynamoDBClient()
	userService := service.NewUserService(dbClient)

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		// Check if the token is valid
		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			sk := claims["id"].(string) // Extract the "sk" claim
			user, _ := userService.ReadUser(sk)
			if user.Pk == "" {
				log.Println("Invalid user")
				return
			}
		} else {
			log.Println("Invalid token")
		}

		if err != nil || !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AuthMiddleware checks for a valid API key in the Authorization header
func AuthSimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" || len(auth) <= 7 || auth[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header"})
			c.Abort()
			return
		}

		token := auth[7:] // Remove "Bearer " prefix

		// Check if token matches any valid API key
		isValid := false
		for _, apiKey := range validAPIKeys {
			if token == apiKey {
				isValid = true
				break
			}
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
