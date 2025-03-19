package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type User struct {
	Email       string `json:"email"`
	Username    string `json:"name"`
	Password    string `json:"password"`
	Conformpass string `json:"conformpass"`
}
type Loginuser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtSecret = []byte("your-secret-key")

func Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Tempuser User
		err := json.NewDecoder(r.Body).Decode(&Tempuser)
		if err != nil {
			http.Error(w, "Invalid Request ", http.StatusBadRequest)
			return
		}
		var exists bool

		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", Tempuser.Email).Scan(&exists)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		if exists {
			http.Error(w, "Email already registered", http.StatusConflict)
			return
		}
		if r.Method == http.MethodPost {
			query := "INSERT INTO users (email,name,password) VALUES ($1,$2,$3)"
			_, err = db.Exec(query, Tempuser.Email, Tempuser.Username, Tempuser.Password)
			if err != nil {
				fmt.Println("Error %v", err)
				http.Error(w, "Error creating user", http.StatusInternalServerError)
				return
			}
			w.Header().Set("X-CustomHeader", "201")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("User registered successfully"))
		}

	}
}

// json.Unmarshal and json.NewDecoder

// json.Unmarshla we have to write an extra code where we have to conver r.Body to bytes data and for that one we have to implement io.ReadAll to convert it.
// but in json.NewDecoder we dont have to implement this thing. just direct implementaion.

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Newuser Loginuser
		err := json.NewDecoder(r.Body).Decode(&Newuser)
		if err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}
		var dbPassword string
		query := "SELECT password from users WHERE email=$1"
		//db.QueryRow query only a single row!
		err = db.QueryRow(query, Newuser.Email).Scan(&dbPassword)
		if err != nil {
			http.Error(w, "Password or email is incorrect", http.StatusUnauthorized)
		}
		if Newuser.Password != dbPassword {
			http.Error(w, "password incorrect ", http.StatusUnauthorized)
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": Newuser.Email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenstring, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Failed to generate token string", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{token:%s}`, tokenstring)))
	}
}
