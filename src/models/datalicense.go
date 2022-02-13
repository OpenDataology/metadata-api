package models

import (
	"fmt"

	"github.com/dataset-license/portal-backend/src/database"
	"github.com/dataset-license/portal-backend/src/utils"
	"github.com/spf13/cast"
)

type Datalicense struct {
	Id                          int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	LicenseUuid                 string `gorm:"type:TEXT" json:"license_uuid,omitempty"`
	LicenseName                 string `gorm:"type:TEXT" json:"license_name,omitempty"`
	LicenseType                 string `gorm:"type:TEXT" json:"license_type,omitempty"`
	OsiApproved                 string `gorm:"type:TEXT" json:"osi_approved,omitempty"`
	ShortIdentifier             string `gorm:"type:TEXT" json:"short_identifier,omitempty"`
	DataAccessRights            string `gorm:"type:TEXT" json:"data_access_rights,omitempty"`
	DataAccessObligations       string `gorm:"type:TEXT" json:"data_access_obligations,omitempty"`
	DataAccessLimitations       string `gorm:"type:TEXT" json:"data_access_limitations,omitempty"`
	DataTaggingRights           string `gorm:"type:TEXT" json:"data_tagging_rights,omitempty"`
	DataTaggingObligations      string `gorm:"type:TEXT" json:"data_tagging_obligations,omitempty"`
	DataTaggingLimitations      string `gorm:"type:TEXT" json:"data_tagging_limitations,omitempty"`
	DataDistributeRights        string `gorm:"type:TEXT" json:"data_distribute_rights,omitempty"`
	DataDistributeObligations   string `gorm:"type:TEXT" json:"data_distribute_obligations,omitempty"`
	DataDistributeLimitations   string `gorm:"type:TEXT" json:"data_distribute_limitations,omitempty"`
	DataNetworkRights           string `gorm:"type:TEXT" json:"data_network_rights,omitempty"`
	DataNetworkObligations      string `gorm:"type:TEXT" json:"data_network_obligations,omitempty"`
	DataNetworkLimitations      string `gorm:"type:TEXT" json:"data_network_limitations,omitempty"`
	DataRepresentRights         string `gorm:"type:TEXT" json:"data_represent_rights,omitempty"`
	DataRepresentObligations    string `gorm:"type:TEXT" json:"data_represent_obligations,omitempty"`
	DataRepresentLimitations    string `gorm:"type:TEXT" json:"data_represent_limitations,omitempty"`
	DataModificationRights      string `gorm:"type:TEXT" json:"data_modification_rights,omitempty"`
	DataModificationObligations string `gorm:"type:TEXT" json:"data_modification_obligations,omitempty"`
	DataModificationLimitations string `gorm:"type:TEXT" json:"data_modification_limitations,omitempty"`
	ModelBenchmarkRights        string `gorm:"type:TEXT" json:"model_benchmark_rights,omitempty"`
	ModelBenchmarkObligations   string `gorm:"type:TEXT" json:"model_benchmark_obligations,omitempty"`
	ModelBenchmarkLimitations   string `gorm:"type:TEXT" json:"model_benchmark_limitations,omitempty"`
	ModelResearchRights         string `gorm:"type:TEXT" json:"model_research_rights,omitempty"`
	ModelResearchObligations    string `gorm:"type:TEXT" json:"model_research_obligations,omitempty"`
	ModelResearchLimitations    string `gorm:"type:TEXT" json:"model_research_limitations,omitempty"`
	ModelPublishRights          string `gorm:"type:TEXT" json:"model_publish_rights,omitempty"`
	ModelPublishObligations     string `gorm:"type:TEXT" json:"model_publish_obligations,omitempty"`
	ModelPublishLimitations     string `gorm:"type:TEXT" json:"model_publish_limitations,omitempty"`
	ModelInternalRights         string `gorm:"type:TEXT" json:"model_internal_rights,omitempty"`
	ModelInternalObligations    string `gorm:"type:TEXT" json:"model_internal_obligations,omitempty"`
	ModelInternalLimitations    string `gorm:"type:TEXT" json:"model_internal_limitations,omitempty"`
	ModelOutputComRights        string `gorm:"type:TEXT" json:"model_output_com_rights,omitempty"`
	ModelOutputComObligations   string `gorm:"type:TEXT" json:"model_output_com_obligations,omitempty"`
	ModelOutputComLimitations   string `gorm:"type:TEXT" json:"model_output_com_limitations,omitempty"`
	ModelComRights              string `gorm:"type:TEXT" json:"model_com_rights,omitempty"`
	ModelComObligations         string `gorm:"type:TEXT" json:"model_com_obligations,omitempty"`
	ModelComLimitations         string `gorm:"type:TEXT" json:"model_com_limitations,omitempty"`
	ModelRevRights              string `gorm:"type:TEXT" json:"model_rev_rights,omitempty"`
	ModelRevObligations         string `gorm:"type:TEXT" json:"model_rev_obligations,omitempty"`
	ModelRevLimitations         string `gorm:"type:TEXT" json:"model_rev_limitations,omitempty"`
	Credit                      string `gorm:"type:TEXT" json:"credit,omitempty"`
	ValidityPeriod              int    `gorm:"type:int" json:"validity_period,omitempty"`
	Liability                   string `gorm:"type:TEXT" json:"liability,omitempty"`
	Designated                  string `gorm:"type:TEXT" json:"designated,omitempty"`
	Additional                  string `gorm:"type:TEXT" json:"additional,omitempty"`
	Remark                      string `gorm:"type:TEXT" json:"remark,omitempty"`
	Available                   int    `gorm:"type:int" json:"available,omitempty"`
}

type LicenseBasic struct {
	Id              int    `json:"id"`
	LicenseUuid     string `json:"license_uuid,omitempty"`
	LicenseName     string `json:"license_name,omitempty"`
	LicenseType     string `json:"license_type,omitempty"`
	OsiApproved     string `json:"osi_approved,omitempty"`
	ShortIdentifier string `json:"short_identifier,omitempty"`
}

type LicenseData struct {
	Can        LicenseDataCan        `json:"can,omitempty"`
	Cannot     LicenseDataCannot     `json:"cannot,omitempty"`
	Obligation LicenseDataObligation `json:"obligation,omitempty"`
	Limitation LicenseDataLimitation `json:"limitation,omitempty"`
}

type LicenseDataCan struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseDataCannot struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseDataObligation struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseDataLimitation struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseModel struct {
	Can        LicenseDataCan        `json:"can,omitempty"`
	Cannot     LicenseDataCannot     `json:"cannot,omitempty"`
	Obligation LicenseDataObligation `json:"obligation,omitempty"`
	Limitation LicenseDataLimitation `json:"limitation,omitempty"`
}

type LicenseModelCan struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseModelCannot struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseModelObligation struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseModelLimitation struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
}

type LicenseOther struct {
	Id       int    `json:"id"`
	Property string `json:"property,omitempty"`
	Value    string `json:"value,omitempty"`
}

func GetDatalicensesByPage(p *utils.Pagination) (Datalicenses []Datalicense, err error) {
	err = database.DB.Model(&Datalicense{}).Scopes(p.GormPaginate()).Find(&Datalicenses).Error
	if err != nil {
		return nil, err
	}
	var total int64
	database.DB.Model(&Datalicense{}).Count(&total)
	p.Total = cast.ToInt(total)
	return
}

func GetDatalicenseBasicByID(id int) (Licensebasic *LicenseBasic, err error) {
	var DatalicenseA Datalicense
	Licensebasic = new(LicenseBasic)
	err = database.DB.Model(&Datalicense{}).First(&DatalicenseA, id).Error
	if err != nil {
		return nil, err
	}
	fmt.Print(DatalicenseA.Id)
	Licensebasic.Id = DatalicenseA.Id
	Licensebasic.LicenseName = DatalicenseA.LicenseName
	Licensebasic.LicenseType = DatalicenseA.LicenseType
	Licensebasic.LicenseUuid = DatalicenseA.LicenseUuid
	Licensebasic.OsiApproved = DatalicenseA.OsiApproved
	Licensebasic.ShortIdentifier = DatalicenseA.ShortIdentifier
	return
}
