/*
router.go

Copyright (c) 2022 The OpenDataology Authors 
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/



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
		basic := new(controllers.BasicInfo)
		v1.GET("/data-license", basic.ListDatasetLicenses)
		v1.GET("/dataset", basic.ListDatasets)
		v1.GET("/get_license_index", basic.GetLicenseIndex)
		v1.GET("/get_dataset_index", basic.GetDatasetIndex)
		v1.GET("/get_license_basic_by_id", basic.GetLicenseBasicById)
		v1.GET("/get_license_basic_by_name", basic.GetLicenseBasicByName)
		v1.GET("/search_license_basic_by_name", basic.SearchLicenseBasicByName)
		v1.GET("/get_license_data_by_id", basic.GetLicenseDataById)
		v1.GET("/get_license_model_by_id", basic.GetLicenseModelById)
		v1.GET("/get_license_other_by_id", basic.GetLicenseOtherById)
		v1.GET("/get_dataset_by_id", basic.GetDatasetById)
		v1.GET("/get_dataset_by_name", basic.GetDatasetByName)
		v1.GET("/search_dataset_by_name", basic.SearchDatasetByName)
		v1.POST("/sign_up", basic.SetSignup)
		v1.POST("/sign_in", basic.GetSignin)
		v1.POST("/set_license", basic.SetLicense)
	}
	return router

}
