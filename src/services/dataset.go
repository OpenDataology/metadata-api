package service

import (
	"net/http"

	"github.com/dataset-license/portal-backend/src/models"
	"github.com/gin-gonic/gin"
)

func GetDatasetIDService(c *gin.Context, id int) (h gin.H) {
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

func GetDatasetNameService(c *gin.Context, name string) (h gin.H) {
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
