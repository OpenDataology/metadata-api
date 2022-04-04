package service

import (
	"net/http"

	"github.com/dataset-license/portal-backend/src/models"
	"github.com/gin-gonic/gin"
)

func SetSignUpService(c *gin.Context, account string, password string, verification string, token string) (h gin.H) {
	user, err := models.SetSignup(account, password, verification)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &user,
	}
	return res
}

func GetSignInService(c *gin.Context, account string, password string, verification string, token string) (h gin.H) {
	user, err := models.GetSignin(account, password, verification)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &user,
	}
	return res
}
