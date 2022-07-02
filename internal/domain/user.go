package domain

type User struct {
	Name  string `json:"name"`
	//Email string `json:"email"`
	Password string `json:"passwrod"`
}

type LoginForm struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}