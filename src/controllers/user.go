/*
user.go

Copyright (c) 2022 The OpenDataology Authors 
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/



package controllers

import (
	"net/http"

	service "github.com/dataset-license/portal-backend/src/services"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) SetSignup(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	verification := c.PostForm("verification")
	token := c.PostForm("token")

	res := service.SetSignUpService(c, account, password, verification, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetSignin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	verification := c.PostForm("verification")
	token := c.PostForm("token")

	res := service.GetSignInService(c, account, password, verification, token)
	a.JsonSuccess(c, http.StatusOK, res)
}
