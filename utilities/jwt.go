package utilities

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing the JWT (keep this secret)
var jwtSecret = []byte(os.Getenv("JWTSECRET"))

// GenerateJWT generates a new JWT for a given user ID
func GenerateJWT(userID string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Define claims (token payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates and parses a JWT token
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Extract and return the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
