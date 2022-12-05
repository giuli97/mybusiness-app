package models

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password,omitempty"`
}
