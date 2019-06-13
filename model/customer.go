package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Users struct {
	u *User `json:"Users"`
}
