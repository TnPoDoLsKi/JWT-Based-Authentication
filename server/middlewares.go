package server

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	jwtMiddle := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, err error) {
			return []byte(os.Getenv("JWT-KEY")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddle.Handler(next)
}
