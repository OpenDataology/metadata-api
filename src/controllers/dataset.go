package controllers

import (
	"net/http"

	service "github.com/dataset-license/portal-backend/src/services"
	"github.com/dataset-license/portal-backend/src/utils"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) ListDatasets(c *gin.Context) {
	p := utils.NewPagination(c)
	res := service.GetDatalicensesService(c, p)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetDatasetById(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatasetIDService(c, id)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetDatasetByName(c *gin.Context) {
	name := c.Query("name")
	res := service.GetDatasetNameService(c, name)
	a.JsonSuccess(c, http.StatusOK, res)
}
