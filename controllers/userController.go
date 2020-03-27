package controllers

import (
	"JWT-Based-Authentication/models"
	"JWT-Based-Authentication/utils"
	"encoding/json"
	"errors"
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
		utils.ErrorResponse(w, http.StatusInternalServerError, 1000, err)
		return
	}

	err = json.Unmarshal(b, &requestUser)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, 1001, err)
		return
	}

	err = databaseUser.FindUserByEmail(requestUser.Email)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, 1002, err)
		return
	}

	if requestUser.Email != databaseUser.Email {
		utils.ErrorResponse(w, http.StatusBadRequest, 1003, errors.New("Email or Password invalide"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(requestUser.Password))

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, 1003, errors.New("Email or Password invalide"))
		return
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
		utils.ErrorResponse(w, http.StatusBadRequest, 1004, err)

		return
	}

	err = json.Unmarshal(body, &u)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, 1005, err)
		return
	}

	if bol, err := utils.EmailValidation(u.Email); !bol {
		utils.ErrorResponse(w, http.StatusBadRequest, 1006, err)
		return
	}

	if bol, err := utils.PasswordValidation(u.Password); !bol {
		utils.ErrorResponse(w, http.StatusBadRequest, 1007, err)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, 1008, err)
		return
	}

	u.Password = string(hashedPassword)

	err = u.SaveUser()

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, 1008, err)
		return
	}

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
