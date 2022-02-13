package service

import (
	"net/http"

	"github.com/dataset-license/portal-backend/src/models"
	"github.com/dataset-license/portal-backend/src/utils"

	"github.com/gin-gonic/gin"
)

func GetDatalicensesService(c *gin.Context, p *utils.Pagination) (h gin.H) {
	datalicenses, err := models.GetDatalicensesByPage(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"pageNum":  p.Page,
		"pageSize": p.Size,
		"totalNum": p.Total,
		"data":     datalicenses,
	}
	return res
}

func GetDatalicensebasicService(c *gin.Context, id int) (h gin.H) {
	datalicensebasic, err := models.GetDatalicenseBasicByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datalicensebasic,
	}
	return res
}
