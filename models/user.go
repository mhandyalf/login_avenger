package models

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Role       string `json:"role"`
}
