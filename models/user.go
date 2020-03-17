package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Login_Name string `json:"login_name"`
	Sex        string `json:"sex"`
	Birthday   string `json:"birthday"`
}
