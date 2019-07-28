package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/* Set up a global string for secret */
var signingKey = []byte("supersecret")

// GetToken creates a 24-hour JWT with the given name and admin status as claims
func GetToken(isAdmin bool, name string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = isAdmin
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with secret */
	tokenString, err = token.SignedString(signingKey)
	return tokenString, nil
}

// LogIn gets name and admin status from headers, create JWT, return JWT as JSON
func LogIn(w http.ResponseWriter, r *http.Request) {
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
