/*
basic.go

Copyright (c) 2022 The OpenDataology Authors 
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/



package controllers

import (
	"github.com/gin-gonic/gin"
)

type Basic struct {
}

type BasicInfo struct {
	Basic
}

func (basic *Basic) JsonSuccess(c *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	c.JSON(status, h)
}

func (basic *Basic) JsonFail(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  "fail",
		"message": message,
	})
}
