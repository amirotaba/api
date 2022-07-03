package domain

type LoginForm struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}