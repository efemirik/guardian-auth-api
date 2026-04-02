package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RequireAuth acts as a security wall, only allowing requests with valid JWTs
func RequireAuth(c *gin.Context) {
	// 1. Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	// 2. Format should be "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format. Expected 'Bearer <token>'"})
		c.Abort()
		return
	}

	tokenString := parts[1]
	secret := os.Getenv("JWT_SECRET")

	// 3. Parse and validate the token signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or manipulated token"})
		c.Abort()
		return
	}

	// 4. Extract claims and check expiration date
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired, please login again"})
			c.Abort()
			return
		}
		// Attach user ID to the request context for subsequent handlers
		c.Set("userID", claims["sub"])
	}

	// Token is valid, proceed to the requested route
	c.Next()
}