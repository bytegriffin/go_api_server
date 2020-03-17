package dao

import (
	"go_api_server/models"
)

func CreateUser(user *models.User) (err error) {
	err = DB.Create(&user).Error
	return
}

func GetUserList() (userList []*models.User, err error) {
	if err := DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetOneUser(id string) (user *models.User, err error) {
	if err = DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateOneUser(user *models.User) (err error) {
	err = DB.Save(user).Error
	return
}

func DeleteOneUser(id string) (err error) {
	err = DB.Where("id?", id).Delete(&models.User{}).Error
	return
}
