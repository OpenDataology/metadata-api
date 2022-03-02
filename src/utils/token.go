package utils

import (
	"github.com/dataset-license/portal-backend/src/config"
	"github.com/dataset-license/portal-backend/src/database"
)

type User struct {
	Id    int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	User  string `gorm:"type:TEXT" json:"user,omitempty"`
	Token string `gorm:"type:TEXT" json:"token,omitempty"`
}

func GetToken() (isAuth bool, err error) {
	config := config.Get()
	var _User User
	err = database.DB.Model(&User{}).First(&_User).Error
	if err != nil {
		return false, err
	}
	if _User.Token == config.TOKEN {
		return true, nil
	} else {
		return false, nil
	}
}
