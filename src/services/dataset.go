package service

import (
	"net/http"

	"github.com/dataset-license/portal-backend/src/models"
	"github.com/dataset-license/portal-backend/src/utils"
	"github.com/gin-gonic/gin"
)

func GetDatasetsService(c *gin.Context, p *utils.Pagination, token string) (h gin.H) {
	datasets, err := models.GetDatasetsByPage(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"pageNum":  p.Page,
		"pageSize": p.Size,
		"totalNum": p.Total,
		"data":     datasets,
	}
	return res
}

func GetDatasetIDService(c *gin.Context, id int, token string) (h gin.H) {
	datasetid, err := models.GetDatasetByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datasetid,
	}
	return res
}

func GetDatasetNameService(c *gin.Context, name string, token string) (h gin.H) {
	datasetname, err := models.GetDatasetByName(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		return
	}
	res := gin.H{
		"data": &datasetname,
	}
	return res
}
