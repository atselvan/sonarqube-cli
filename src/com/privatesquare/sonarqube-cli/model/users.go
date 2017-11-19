package model

type AuthUser struct {
	Username string
	Password string
	Valid bool `json:"valid"`
}

type SonarUser struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
