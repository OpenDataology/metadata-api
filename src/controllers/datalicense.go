/*
datalicense.go

Copyright (c) 2022 The OpenDataology Authors 
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/



package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dataset-license/portal-backend/src/models"
	service "github.com/dataset-license/portal-backend/src/services"
	"github.com/dataset-license/portal-backend/src/utils"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) ListDatasetLicenses(c *gin.Context) {
	token := c.Query("token")
	tp := [4]string{"all", "Data License", "Data-Specific License", "DataSource Terms of Use"}
	p := utils.NewPagination(c)
	int_tp := cast.ToInt(c.Query("type"))
	t := tp[int_tp]
	res := service.GetDatalicensesService(c, p, t, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseBasicById(c *gin.Context) {
	token := c.Query("token")
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicensebasicService(c, id, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseBasicByName(c *gin.Context) {
	token := c.Query("token")
	name := c.Query("name")
	res := service.GetDatalicensebasicnameService(c, name, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) SearchLicenseBasicByName(c *gin.Context) {
	token := c.Query("token")
	tp := [4]string{"all", "Data License", "Data-Specific License", "DataSource Terms of Use"}
	name := c.Query("name")
	int_tp := cast.ToInt(c.Query("type"))
	t := tp[int_tp]
	res := service.SearchDatalicensebasicnameService(c, name, t, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseDataById(c *gin.Context) {
	token := c.Query("token")
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicenseDataService(c, id, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseModelById(c *gin.Context) {
	token := c.Query("token")
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicenseModelService(c, id, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseOtherById(c *gin.Context) {
	token := c.Query("token")
	id := cast.ToInt(c.Query("id"))
	res := service.GetDatalicenseOtherService(c, id, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) GetLicenseIndex(c *gin.Context) {
	token := c.Query("token")
	tp := [4]string{"all", "Data License", "Data-Specific License", "DataSource Terms of Use"}
	int_tp := cast.ToInt(c.Query("type"))
	t := tp[int_tp]
	res := service.GetDatalicensesIndexService(c, t, token)
	a.JsonSuccess(c, http.StatusOK, res)
}

func (a *BasicInfo) SetLicense(c *gin.Context) {
	license := c.PostForm("license")
	token := c.PostForm("token")
	var data models.LicenseUpload

	if err := json.Unmarshal([]byte(license), &data); err != nil {
		a.JsonSuccess(c, http.StatusOK, nil)
	} else {
		res := service.SetLicense(c, data, token)
		a.JsonSuccess(c, http.StatusOK, res)
	}

}
