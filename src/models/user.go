package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/dataset-license/portal-backend/src/database"
)

type User struct {
	Id           int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	Account      string `gorm:"type:TEXT" json:"account,omitempty"`
	Password     string `gorm:"type:TEXT" json:"password,omitempty"`
	Verification string `gorm:"type:TEXT" json:"verification,omitempty"`
}

func SetSignup(account string, password string, verification string) (user *User, err error) {
	var count int64

	database.DB.Model(&User{}).Where("account = ?", account).Count(&count)
	if count > 0 {
		return nil, fmt.Errorf("account already exist!")
	}
	password_byte := []byte(password)
	md5Ctx := md5.New()
	md5Ctx.Write(password_byte)
	cipherStr := md5Ctx.Sum(nil)
	user = &User{
		Account:      account,
		Password:     hex.EncodeToString(cipherStr),
		Verification: verification,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return
}

func GetSignin(account string, password string, verification string) (user *User, err error) {
	var count int64
	database.DB.Model(&User{}).Where("account = ?", account).Count(&count)
	if count == 0 {
		return
	}
	err = database.DB.Model(&User{}).Where("account = ?", account).First(&user).Error
	password_byte := []byte(password)
	md5Ctx := md5.New()
	md5Ctx.Write(password_byte)
	cipherStr := md5Ctx.Sum(nil)
	strCipher := hex.EncodeToString(cipherStr)
	if user.Password != strCipher {
		return nil, err
	}
	return
}
