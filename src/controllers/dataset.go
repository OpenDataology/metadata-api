package controllers

import (
	"net/http"

	service "github.com/dataset-license/portal-backend/src/services"
	"github.com/dataset-license/portal-backend/src/utils"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) ListDatasets(c *gin.Context) {
	p := utils.NewPagination(c)
	res := service.GetDatalicensesService(c, p)
	a.JsonSuccess(c, http.StatusOK, res)
}
