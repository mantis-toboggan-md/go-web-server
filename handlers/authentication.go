package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/* Set up a global string for secret */
var signingKey = []byte("supersecret")

// GetToken creates a 24-hour JWT with the given name and admin status as claims
func GetToken(isAdmin bool, name string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": isAdmin,
		"name":  name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	/* Sign the token with secret */
	tokenString, err = token.SignedString(signingKey)
	return tokenString, nil
}

// check that the JWT is using the signing method expected
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return signingKey, nil
}

// ValidateToken takes in a token, parses, returns bool
func ValidateToken(tokenString string) (isValid bool, err error) {
	if token, err := jwt.Parse(tokenString, keyFunc); err != nil {
		return false, err
	} else {
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return true, nil
		}
		return false, nil
	}

}

// NeedsToken middleware checks validity of tokenString in req header then continues or sends 403 forbidden
func NeedsToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("authorization")
		isValid, err := ValidateToken(tokenString)
		if err != nil || !isValid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// LogIn gets name and admin status from headers, create JWT, return JWT as JSON
func LogIn(w http.ResponseWriter, r *http.Request) {
	// TODO: add some kind of authentication here
	name := r.Header.Get("name")
	isAdmin := r.Header.Get("isAdmin") == "true"
	tokenString, err := GetToken(isAdmin, name)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	data := map[string]string{"token": tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
