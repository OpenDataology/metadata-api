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
	token := c.Query("token")
	res := service.GetDatasetsService(c, p, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetDatasetById(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	token := c.Query("token")
	res := service.GetDatasetIDService(c, id, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetDatasetByName(c *gin.Context) {
	name := c.Query("name")
	token := c.Query("token")
	res := service.GetDatasetNameService(c, name, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) SearchDatasetByName(c *gin.Context) {
	name := c.Query("name")
	token := c.Query("token")
	res := service.SearchDatasetNameService(c, name, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetDatasetIndex(c *gin.Context) {
	token := c.Query("token")
	res := service.GetDatasetIndexService(c, token)
	a.JsonSuccess(c, http.StatusOK, res)
}
