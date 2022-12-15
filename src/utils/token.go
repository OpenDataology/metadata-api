/*
token.go

Copyright (c) 2022 The OpenDataology Authors 
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/



package utils

import (
	"github.com/dataset-license/portal-backend/src/config"
)

type User struct {
	Id    int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	User  string `gorm:"type:TEXT" json:"user,omitempty"`
	Token string `gorm:"type:TEXT" json:"token,omitempty"`
}

func GetToken(token string) (isAuth bool, err error) {
	config := config.Get()
	if token == config.TOKEN {
		return true, nil
	} else {
		return false, nil
	}
}
