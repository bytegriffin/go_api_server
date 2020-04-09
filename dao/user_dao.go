package dao

import (
	"go_api_server/models"
	"log"
)

func InsertUser(user models.User) (err error) {
	sql := "insert into users(login_name) values (?)"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Fatal("insert failed. err: %v\n", err)
	}
	defer stmt.Close()
	ret, err := stmt.Exec(user.Login_Name)
	if err != nil {
		log.Fatal("insert failed. err: %v\n", err)
	}
	insertId, err := ret.LastInsertId()
	if err != nil {
		log.Fatal("get insert id %s failed. err: %v\n", insertId, err)
	}
	return
}

func GetUserList() (userList []models.User, err error) {
	sql := "select id, name, login_name, sex, birthday from users"
	rows, err := DB.Query(sql)
	if err != nil {
		log.Fatal("Query %s failed, err: %v\n", sql, err)
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Login_Name, &u.Sex, &u.Birthday)
		if err != nil {
			log.Fatal("Scan failed, err: %v\n", err)
		}
		userList = append(userList, u)
	}
	return
}

func GetOneUser(id string) (user models.User, err error) {
	sql := "select id,name,login_name,birthday,sex from users where id=?"
	row := DB.QueryRow(sql, id)
	err = row.Scan(&user.ID, &user.Name, &user.Login_Name, &user.Sex, &user.Birthday)
	if err != nil {
		log.Fatal("Scan %s failed, err: %v\n", sql, err)
	}
	return
}

func UpdateOneUser(user models.User) (err error) {
	sql := "update users set name=? where id = ?"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Fatal("update failed. err：%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Login_Name, user.ID)
	if err != nil {
		log.Fatal("update failed. err：%v\n", err)
		return
	}
	return
}

func DeleteOneUser(id string) (err error) {
	sql := "delete from users where id = ?"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Fatal("delete failed, err: %v ", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal("delete failed. err：%v ", err)
		return
	}
	return
}
