package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func New(username, password string, roleID int) *User {

	return &User{
		Username: username,
		Password: password,
		Role:     roleID,
	}
}
