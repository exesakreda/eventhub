package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationCredentials struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
