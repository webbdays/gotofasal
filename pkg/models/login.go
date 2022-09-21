package models

type SignUpData struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=15"`
}

type SignInData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=15"`
}
