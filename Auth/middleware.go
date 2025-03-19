package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// To get the Authorization Header we need to do request.Header.Get function which will
		// eventually gives you the Authorization Header.
		token := request.Header.Get("Authorization")
		if token == "" {
			http.Error(writer, "MissingAuthorizationheader", http.StatusUnauthorized)
		}
		newtoken := strings.TrimPrefix(token, "Bearer ")
		tokenn, err := jwt.Parse(newtoken, func(token *jwt.Token) (interface{}, error) {
			// Here you can return any thing , and its going to be an  recognized by an interface.
			return jwtSecret, nil
		})
		if err != nil || !tokenn.Valid {
			fmt.Printf("Value %v", err)
			http.Error(writer, "Error while parsing the tokenn", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
