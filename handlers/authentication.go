package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mantis_toboggan_md/go_test/mydb"
	"golang.org/x/crypto/bcrypt"
)

/* Set up a global string for secret */
var signingKey = []byte("supersecret")

// GetToken creates a 24-hour JWT with the given name admin status and id as claims
func GetToken(isAdmin bool, name string, id int64) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": isAdmin,
		"name":  name,
		"id":    id,
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

// return generic ok message for token route
func Validated(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "token validated")
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

/*
* CreateAccount gets User from req.body and adds to db
 */
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// get user from req body
	var userReq mydb.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userReq); err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.MinCost)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "unknown server error")
		return
	}

	//make User struct to add to db
	userDb := mydb.User{
		Name:     userReq.Name,
		IsAdmin:  userReq.IsAdmin,
		Password: string(hash),
	}
	//open db connection, insert user, defer close
	db, err := mydb.PingDB("postgres", mydb.PgConnection)
	_, err = mydb.InsertOneUser(db, userDb)
	defer db.Close()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "user created successfully")
	return

}

/* LogIn gets user from req body
 * create JWT, return JWT as JSON
 */
func LogIn(w http.ResponseWriter, r *http.Request) {
	// get user from req body
	var userReq mydb.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userReq); err != nil {
		//return err
	}

	// get user from db
	db, err := mydb.PingDB("postgres", mydb.PgConnection)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal server error")
		return
	}
	var userDB mydb.User

	userDB, err = mydb.GetOneUser(db, userReq.Name)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, err.Error())
		return
	}

	defer db.Close()
	dbPwd := []byte(userDB.Password)
	userPwd := []byte(userReq.Password)

	// return http error if bad pwd
	if err := bcrypt.CompareHashAndPassword(dbPwd, userPwd); err != nil {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Invalid password")
		return
	}

	// otherwise, make token and return with user data

	tokenString, err := GetToken(userDB.IsAdmin, userDB.Name, userDB.Id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	token := map[string]string{"token": tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
	//return nil
}
