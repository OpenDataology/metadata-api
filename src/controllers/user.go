package controllers

import (
	"net/http"

	service "github.com/dataset-license/portal-backend/src/services"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) SetSignup(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	verification := c.Query("verification")
	token := c.Query("token")

	res := service.SetSignUpService(c, account, password, verification, token)
	if res == nil{
		a.JsonFail(c, http.StatusOK, "user already exist")
	}
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetSignin(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	verification := c.Query("verification")
	token := c.Query("token")

	res := service.GetSignInService(c, account, password, verification, token)
	a.JsonSuccess(c, http.StatusOK, res)
}
