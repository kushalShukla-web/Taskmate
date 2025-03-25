package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Here we are checking for cros origin request and if that request is there then
		// send 200 response to that.
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}

		// To get the Authorization Header we need to do request.Header.Get function which will
		// eventually gives you the Authorization Header.
		token := request.Header.Get("Authorization")
		if token == "" {
			http.Error(writer, "MissingAuthorizationheader", http.StatusUnauthorized)
		}
		newtoken := strings.TrimPrefix(token, "Bearer ")
		if len(strings.Split(token, ".")) != 3 {
			fmt.Printf("Error while spliting")
			return
		}
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
