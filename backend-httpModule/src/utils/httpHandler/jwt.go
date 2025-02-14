package httpHandler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleCookie = http.Cookie{
	Name:     "auth",
	Value:    "", // Set an empty value
	Path:     "/",
	MaxAge:   -1, // Set MaxAge to a negative value to immediately expire the cookie
	HttpOnly: true,
	Secure:   true,
	SameSite: http.SameSiteLaxMode,
}

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 2).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key used to sign the token
		return secretKey, nil
	})

	if err != nil {
		return err
	}



	// Check if the token is valid
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// Check for expiration
	if expClaim, ok := token.Claims.(jwt.MapClaims)["exp"]; ok {
		expTime, ok := expClaim.(float64)
		if !ok {
			return fmt.Errorf("invalid expiration claim")
		}
		// Convert expiration time from float64 to time.Time
		expiration := time.Unix(int64(expTime), 0)
		if time.Now().After(expiration) {
			return fmt.Errorf("token has expired")
		}
	}

	// Check for not before time
	if nbfClaim, ok := token.Claims.(jwt.MapClaims)["nbf"]; ok {
		nbfTime, ok := nbfClaim.(float64)
		if !ok {
			return fmt.Errorf("invalid not-before claim")
		}
		// Convert not-before time from float64 to time.Time
		notBefore := time.Unix(int64(nbfTime), 0)
		if time.Now().Before(notBefore) {
			return fmt.Errorf("token is not valid yet")
		}
	}



	return nil
}

func GetUserNameAndToken(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key used to sign the token
		return secretKey, nil
	})

	if username, ok := token.Claims.(jwt.MapClaims)["username"]; ok {
		return username.(string)
	}
	return ""
}

// MiddleWere for All
// 1. Use Cookie to retrieve data and verify user
// 2.  added checking Header Auth Check also as if cookie is not Available
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to retrieve the token from the 'Authorization' header
		authHeader := r.Header.Get("Authorization")

		// If 'Authorization' header is present
		if authHeader != "" {
			// Check for Bearer token
			if strings.HasPrefix(authHeader, "Bearer ") {
				token := strings.TrimPrefix(authHeader, "Bearer ")

				// Verify the token
				if err := VerifyToken(token); err != nil {
					// If token verification fails, respond with an error
					http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
					return
				}
			} else {
				// If the header format is not "Bearer <token>"
				http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
				return
			}
		} else {
			// If no Authorization header, check for the 'auth' cookie
			cookie, err := r.Cookie("auth")
			if err == nil {
				// If cookie exists, verify its token
				if err := VerifyToken(cookie.Value); err != nil {
					// If token verification fails, respond with an error
					http.Error(w, "Invalid authentication token in cookie", http.StatusUnauthorized)
					return
				}
			} else {
				// If no 'auth' cookie and no 'Authorization' header, deny access
				http.Error(w, "Missing authentication token", http.StatusUnauthorized)
				return
			}
		}

		// Proceed to the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}
