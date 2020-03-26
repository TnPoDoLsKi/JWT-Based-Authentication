package controllers

import (
	"JWT-Based-Authentication/models"
	"JWT-Based-Authentication/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var b []byte
	var databaseUser, requestUser models.User

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(b, &requestUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	databaseUser.FindUserByEmail(requestUser.Email)

	if requestUser.Email != databaseUser.Email {
		http.Error(w, "Email not found", http.StatusBadRequest)
	}

	err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(requestUser.Password))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokenString, err := GenerateToken(databaseUser.Id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User

	w.Header().Set("content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if bol, err := utils.EmailValidation(u.Email); !bol {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if bol, err := utils.PasswordValidation(u.Password); !bol {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	u.Password = string(hashedPassword)

	u.SaveUser()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created successfully"))
}

func GenerateToken(id uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"expiration": time.Now().Add(time.Hour * time.Duration(5)).Unix(),
		"creation":   time.Now(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT-KEY")))

	return tokenString, err
}
