package main

import (
	"login_avenger/database"
	"login_avenger/handlers"
	"net/http"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	router := handlers.NewRouter(db)

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
