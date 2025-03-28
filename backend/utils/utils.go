package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"

	"time"
	"github.com/dgrijalva/jwt-go"
	// "log"
)

// Secret key used to sign the JWT token (should be kept secure and not hardcoded in production)
var secretKey = []byte("your-secret-key") // You should replace this with an environment variable in production

// GenerateToken generates a JWT token for the user
func GenerateToken(userID uint) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error signing token:", err)
		return "", err
	}

	return tokenString, nil
}

// HashPassword hashes the given password using bcrypt and returns the hashed password or an error.
func HashPassword(password string) (string, error) {
	// Generate a hashed password using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash checks if the provided password matches the stored hashed password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
