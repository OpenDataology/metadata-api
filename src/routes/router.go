package routes

import (
	"github.com/dataset-license/portal-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	v1 := router.Group("/api/v1")
	{
		dataLicense := new(controllers.BasicInfo)
		v1.GET("/data-license", dataLicense.Index)
		v1.GET("/get_license_basic_by_id", dataLicense.GetLicenseBasicById)
		v1.GET("/get_license_data_by_id", dataLicense.GetLicenseDataById)
	}

	return router

}
