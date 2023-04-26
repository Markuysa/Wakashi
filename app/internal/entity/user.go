package entity

type User struct {
	Username string `json:"Username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
