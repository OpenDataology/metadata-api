package routes

import (
	"github.com/dataset-license/portal-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		dataLicense := new(controllers.DataLicense)
		v1.GET("/data-license", dataLicense.Index)
	}

	return router

}
