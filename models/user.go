package models

import "database/sql"

type User struct {
	ID         int            `json:"id"`
	Name       sql.NullString `json:"name"`
	Login_Name string         `json:"login_name"`
	Sex        sql.NullString `json:"sex"`
	Birthday   sql.NullString `json:"birthday"`
}
