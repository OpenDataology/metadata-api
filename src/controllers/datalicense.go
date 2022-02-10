package controllers

import (
	"net/http"

	"github.com/dataset-license/portal-backend/src/database"
	"github.com/dataset-license/portal-backend/src/models"

	"github.com/gin-gonic/gin"
)

type DataLicense struct {
	Basic
}

func (a *DataLicense) Index(c *gin.Context) {
	var datalicenses []models.Datalicense
	database.DB.Find(&datalicenses)
	a.JsonSuccess(c, http.StatusOK, gin.H{"data": datalicenses})
}
