package model

type User struct {
	ID       int   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"-"`
	UserType int    `json:"usertype"`
}
