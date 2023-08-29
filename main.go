package main

import (
	"fmt"
	"login_avenger/api"
	"login_avenger/auth"
	"login_avenger/config"
	"login_avenger/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	cfg := config.NewConfig()
	auth.Initialize(cfg)

	// Menggunakan middleware LoggingMiddleware untuk mencatat log
	router.POST("/register", middleware.LoggingMiddleware(api.RegisterHandler))
	router.POST("/login", middleware.LoggingMiddleware(api.LoginHandler))

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
