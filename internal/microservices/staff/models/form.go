package models

import "time"

//easyjson:json
type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailConfirmationForm struct {
	SecretKey      string    `json:"secret_key"`
	Email          string    `json:"email" validate:"required,email"`
	DateOfCreation time.Time `json:"date_of_creation"`
	IsRegistered   bool      `json:"is_registered"`
}
