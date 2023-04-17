package entity

type User struct {
	Username string
	Password string
	Role     string
}

func New(username, password, role string) *User {

	return &User{
		Username: username,
		Password: password,
		Role:     role,
	}
}
