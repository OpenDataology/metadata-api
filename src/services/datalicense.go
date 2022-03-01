package service

import (
	"net/http"

	"github.com/dataset-license/portal-backend/src/models"
	"github.com/dataset-license/portal-backend/src/utils"

	"github.com/gin-gonic/gin"
)

func GetDatalicensesService(c *gin.Context, p *utils.Pagination, t string) (h gin.H) {
	datalicenses, err := models.GetDatalicensesByPage(p, t)
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

func GetDatalicensebasicnameService(c *gin.Context, name string) (h gin.H) {
	datalicensebasic, err := models.GetDatalicenseBasicByName(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datalicensebasic,
	}
	return res
}

func GetDatalicenseDataService(c *gin.Context, id int) (h gin.H) {
	datalicensedata, err := models.GetDatalicenseDataByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datalicensedata,
	}
	return res
}

func GetDatalicenseModelService(c *gin.Context, id int) (h gin.H) {
	datalicensemodel, err := models.GetDatalicenseModelByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datalicensemodel,
	}
	return res
}

func GetDatalicenseOtherService(c *gin.Context, id int) (h gin.H) {
	datalicenseother, err := models.GetDatalicenseOtherByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datalicenseother,
	}
	return res
}
