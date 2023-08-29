package api

import (
	"encoding/json"
	"login_avenger/auth"
	"login_avenger/validators"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newUser auth.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidateUser(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth.RegisterUser(newUser)

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := auth.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func InitializeRoutes(router *httprouter.Router) {
	router.POST("/register", RegisterHandler)
	router.POST("/login", LoginHandler)
}
