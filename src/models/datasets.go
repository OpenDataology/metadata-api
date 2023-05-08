/*
datasets.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

import (
	"github.com/dataset-license/portal-backend/src/database"
	"github.com/dataset-license/portal-backend/src/utils"
	"github.com/spf13/cast"
)

type Dataset struct {
	Id                   int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	DatasetName          string `gorm:"type:TEXT" json:"dataset_name,omitempty"`
	DatasetVersion       string `gorm:"type:TEXT" json:"dataset_version,omitempty"`
	LicenseId            int    `gorm:"type:int" json:"license_id,omitempty"`
	LicenseName          string `gorm:"type:TEXT" json:"license_name,omitempty"`
	Licensor             string `gorm:"type:TEXT" json:"licensor,omitempty"`
	LicenseFrom          string `gorm:"type:TEXT" json:"license_from,omitempty"`
	LicenseLocation      string `gorm:"type:TEXT" json:"license_location,omitempty"`
	TaskType             string `gorm:"type:TEXT" json:"task_type,omitempty"`
	LicenseContent       string `gorm:"type:TEXT" json:"license_content,omitempty"`
	Origin               string `gorm:"type:TEXT" json:"origin,omitempty"`
	DownloadedOutlet     string `gorm:"type:TEXT" json:"downloaded_outlet,omitempty"`
	OutletLicensed       string `gorm:"type:TEXT" json:"outlet_licensed,omitempty"`
	HashCode             string `gorm:"type:TEXT" json:"hash_code,omitempty"`
	DataSize             string `gorm:"type:TEXT" json:"data_size,omitempty"`
	Format               string `gorm:"type:TEXT" json:"format,omitempty"`
	Description          string `gorm:"type:TEXT" json:"description,omitempty"`
	IsPersonalData       string `gorm:"type:TEXT" json:"is_personal_data,omitempty"`
	IsAdditionalVerify   string `gorm:"type:TEXT" json:"is_additional_verify,omitempty"`
	IsOffensiveContent   string `gorm:"type:TEXT" json:"is_offensive_content,omitempty"`
	CollectionProcess    string `gorm:"type:TEXT" json:"collection_process,omitempty"`
	Restrictions         string `gorm:"type:TEXT" json:"restrictions,omitempty"`
	IsComply             string `gorm:"type:TEXT" json:"is_comply,omitempty"`
	RestrictionNotes     string `gorm:"type:TEXT" json:"restriction_notes,omitempty"`
	DatasetCollectMethod string `gorm:"type:TEXT" json:"dataset_collect_method,omitempty"`
	AdditionalNotes      string `gorm:"type:TEXT" json:"additional_notes,omitempty"`
	Challenges           string `gorm:"type:TEXT" json:"challenges,omitempty"`
	Available            int    `gorm:"type:int" json:"available,omitempty"`
}

type DatasetIndex struct {
	Id          int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	DatasetName string `gorm:"type:TEXT" json:"dataset_name,omitempty"`
}

func GetDatasetsByPage(p *utils.Pagination) (Datasets []Dataset, err error) {
	err = database.DB.Model(&Dataset{}).Scopes(p.GormPaginate()).Find(&Datasets).Error
	if err != nil {
		return nil, err
	}
	var total int64
	database.DB.Model(&Dataset{}).Count(&total)
	p.Total = cast.ToInt(total)
	return
}

func GetDatasetByID(id int) (_Dataset *Dataset, err error) {
	err = database.DB.Model(&Dataset{}).First(&_Dataset, id).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetDatasetByName(name string) (_Dataset *Dataset, err error) {
	err = database.DB.Model(&Dataset{}).Where("dataset_name = ?", name).First(&_Dataset).Error
	if err != nil {
		return nil, err
	}
	return
}

func SearchDatasetByName(name string) (_Dataset []Dataset, err error) {
	err = database.DB.Model(&Dataset{}).Where("dataset_name LIKE ?", "%"+name+"%").Limit(20).Find(&_Dataset).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetDatasetIndex() (Datasetindexes []DatasetIndex, err error) {
	err = database.DB.Model(&Dataset{}).Find(&Datasetindexes).Error
	if err != nil {
		return nil, err
	}
	return
}
