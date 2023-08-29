package auth

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Role       string `json:"role"`
}

var db *sql.DB

func RegisterUser(user User) error {
	// Simpan data user ke database
	_, err := db.Exec("INSERT INTO user_data (email, password, full_name, age, occupation, role) VALUES (?, ?, ?, ?, ?, ?)",
		user.Email, user.Password, user.FullName, user.Age, user.Occupation, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(email, password string) (string, error) {
	var token string
	err := db.QueryRow("SELECT role FROM user_data WHERE email = ? AND password = ?", email, password).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	return token, nil
}
