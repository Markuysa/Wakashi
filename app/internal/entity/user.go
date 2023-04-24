package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func NewUser(username string, password string, role int) *User {
	return &User{Username: username, Password: password, Role: role}
}
