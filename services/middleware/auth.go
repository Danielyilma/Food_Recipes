package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Danielyilma/Food_Recipes/services/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/square/go-jose/v3"
)

// Define a custom key for storing the user ID in the context
type contextKey string

const userIDKey contextKey = "user_id"

// Load the public key for verification (assuming GetJWK is implemented)
func GetPublicKey() (*jose.JSONWebKeySet, error) {
	// Assuming you already have the JWK loading function implemented as GetJWK()
	jwk, err := helpers.GetJWK()
	return &jwk, err
}

// JWT Middleware to verify token and add user data to the context
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
		}
		// Load the public key to verify the token
		jwkSet, err := GetPublicKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading public key"})
			c.Abort()
			return
		}
		// Parse and verify the JWT
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Get the public key from the JWK set based on the key ID in the token
			kid, ok := token.Header["Kid"].(string)
			if !ok {
				return nil, fmt.Errorf("missing key ID")
			}

			var key *jose.JSONWebKey
			for _, jwk := range jwkSet.Keys {
				if jwk.KeyID == kid {
					key = &jwk
					break
				}
			}
			if key == nil {
				return nil, fmt.Errorf("key not found")
			}
			// Return the public key for verification
			return key.Key, nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Token is valid, extract claims and store in context
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			// Extract the user ID from the claims
			userID := claims["user_id"]
			if userID == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}

			// Store the user ID in the context
			c.Set(string(userIDKey), userID)

			// Proceed to the next handler
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}
