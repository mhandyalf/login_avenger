package database

import (
	"database/sql"
	"login_avenger/models"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func InitDB() *DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/avenger")
	if err != nil {
		panic(err)
	}

	return &DB{db}
}

func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password, full_name, age, occupation, role FROM user_data WHERE email = ?"
	user := &models.User{}

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.FullName, &user.Age, &user.Occupation, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // User not found
		}
		return nil, err
	}

	return user, nil
}

func (db *DB) CreateUser(user *models.User) error {
	query := "INSERT INTO user_data (email, password, full_name, age, occupation, role) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, user.Email, user.Password, user.FullName, user.Age, user.Occupation, user.Role)
	if err != nil {
		return err
	}

	return nil
}
