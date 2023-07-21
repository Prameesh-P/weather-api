package model


type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserActivity struct {
	ID       int    `json:"id"`
	Email       string  `json:"email"`
	Temperature float64 `json:"temperature"`
	City        string  `json:"city"`
	Time        string `json:"time"`
}