package main

import (
	"Server/auth"
	"Server/handler"
	"Server/product"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := product.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	db := product.GetDb()
	customhandler := handler.Initializer()

	sm := http.NewServeMux()
	sm.HandleFunc("/signup", auth.Signup(db))
	sm.HandleFunc("/login", auth.Login(db))
	sm.Handle("/", auth.AuthMiddleware(customhandler))
	fmt.Println("Server is running on port 8080...")
	err = http.ListenAndServe(":8080", sm)
	if err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
