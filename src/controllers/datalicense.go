package controllers

import (
	"net/http"

	service "github.com/dataset-license/portal-backend/src/services"
	"github.com/dataset-license/portal-backend/src/utils"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) ListDatasetLicenses(c *gin.Context) {
	p := utils.NewPagination(c)
	res := service.GetDatalicensesService(c, p)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseBasicById(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicensebasicService(c, id)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseBasicByName(c *gin.Context) {
	name := c.Query("name")
	res := service.GetDatalicensebasicnameService(c, name)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseDataById(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicenseDataService(c, id)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseModelById(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicenseModelService(c, id)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseOtherById(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicenseOtherService(c, id)
	a.JsonSuccess(c, http.StatusOK, res)
}
