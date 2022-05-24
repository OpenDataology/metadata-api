package models

import (
	"fmt"
	"reflect"
	"strings"

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

type LicenseIndex struct {
	Id          int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	LicenseName string `gorm:"type:TEXT" json:"license_name,omitempty"`
}

type LicenseBasic struct {
	Id              int    `json:"id,omitempty"`
	LicenseUuid     string `json:"license_uuid,omitempty"`
	LicenseName     string `json:"license_name,omitempty"`
	LicenseType     string `json:"license_type,omitempty"`
	OsiApproved     string `json:"osi_approved,omitempty"`
	ShortIdentifier string `json:"short_identifier,omitempty"`
}

type LicenseData struct {
	Can        []LicenseDataCan        `json:"can,omitempty"`
	Cannot     []LicenseDataCannot     `json:"cannot,omitempty"`
	Obligation []LicenseDataObligation `json:"obligation,omitempty"`
	Limitation []LicenseDataLimitation `json:"limitation,omitempty"`
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
	Can        []LicenseModelCan        `json:"can,omitempty"`
	Cannot     []LicenseModelCannot     `json:"cannot,omitempty"`
	Obligation []LicenseModelObligation `json:"obligation,omitempty"`
	Limitation []LicenseModelLimitation `json:"limitation,omitempty"`
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

type License4Create struct {
	Name             string `json:"name,omitempty"`
	Rights           string `json:"rights,omitempty"`
	Limitations      string `json:"limitations,omitempty"`
	Limitations_Text string `json:"limitations_text,omitempty"`
	Obligations      string `json:"obligations,omitempty"`
	Obligations_Text string `json:"obligations_text,omitempty"`
}

type LicenseData4Create struct {
	Access       License4Create `json:"access,omitempty"`
	Tagging      License4Create `json:"tagging,omitempty"`
	Distribute   License4Create `json:"distribute,omitempty"`
	Network      License4Create `json:"network,omitempty"`
	Represent    License4Create `json:"represent,omitempty"`
	Modification License4Create `json:"modification,omitempty"`
}

type LicenseModel4Create struct {
	Benchmark  License4Create `json:"benchmark,omitempty"`
	Research   License4Create `json:"research,omitempty"`
	Publish    License4Create `json:"publish,omitempty"`
	Internal   License4Create `json:"internal,omitempty"`
	Output_com License4Create `json:"output_com,omitempty"`
	Com        License4Create `json:"com,omitempty"`
	Rev        License4Create `json:"rev,omitempty"`
}

type LicenseOthers4Create struct {
	Liability      string `json:"liability,omitempty"`
	Designated     string `json:"designated,omitempty"`
	Additional     string `json:"additional,omitempty"`
	Credit         string `json:"credit,omitempty"`
	ValidityPeriod int    `json:"validityPeriod,omitempty"`
}

type LicenseUpload struct {
	Data   LicenseData4Create   `json:"data,omitempty"`
	Model  LicenseModel4Create  `json:"model,omitempty"`
	Basics LicenseBasic         `json:"basics,omitempty"`
	Others LicenseOthers4Create `json:"others,omitempty"`
}

func SetDatalicense(license LicenseUpload) (Data_license *Datalicense, err error) {
	var count int64
	name := license.Basics.LicenseName
	database.DB.Model(&Datalicense{}).Where("license_name = ?", name).Count(&count)
	if count > 0 {
		return nil, fmt.Errorf("license_name already exist!")
	}
	Data_license = new(Datalicense)

	//LicenseBasic
	Data_license.LicenseName = license.Basics.LicenseName
	Data_license.LicenseType = license.Basics.LicenseType
	Data_license.OsiApproved = license.Basics.OsiApproved
	Data_license.ShortIdentifier = license.Basics.ShortIdentifier

	//LicenseOthers
	Data_license.OtherAdditional = license.Others.Additional
	Data_license.OtherCredit = license.Others.Credit
	Data_license.OtherDesignated = license.Others.Designated
	Data_license.OtherLiability = license.Others.Liability
	Data_license.OtherValidityPeriod = license.Others.ValidityPeriod

	//License Data Part
	Data_license.DataAccessRights = license.Data.Access.Rights
	Data_license.DataAccessObligations = license.Data.Access.Obligations_Text
	Data_license.DataAccessLimitations = license.Data.Access.Limitations_Text
	Data_license.DataDistributeRights = license.Data.Distribute.Rights
	Data_license.DataDistributeObligations = license.Data.Distribute.Obligations_Text
	Data_license.DataDistributeLimitations = license.Data.Distribute.Limitations_Text
	Data_license.DataModificationRights = license.Data.Modification.Rights
	Data_license.DataModificationObligations = license.Data.Modification.Obligations_Text
	Data_license.DataModificationLimitations = license.Data.Modification.Limitations_Text
	Data_license.DataNetworkRights = license.Data.Network.Rights
	Data_license.DataNetworkObligations = license.Data.Network.Obligations_Text
	Data_license.DataNetworkLimitations = license.Data.Network.Limitations_Text
	Data_license.DataRepresentRights = license.Data.Represent.Rights
	Data_license.DataRepresentObligations = license.Data.Represent.Obligations_Text
	Data_license.DataRepresentLimitations = license.Data.Represent.Limitations_Text
	Data_license.DataTaggingRights = license.Data.Tagging.Rights
	Data_license.DataTaggingObligations = license.Data.Tagging.Obligations_Text
	Data_license.DataTaggingLimitations = license.Data.Tagging.Limitations_Text

	//License Model Part
	Data_license.ModelBenchmarkRights = license.Model.Benchmark.Rights
	Data_license.ModelBenchmarkObligations = license.Model.Benchmark.Obligations_Text
	Data_license.ModelBenchmarkLimitations = license.Model.Benchmark.Limitations_Text
	Data_license.ModelComRights = license.Model.Com.Rights
	Data_license.ModelComObligations = license.Model.Com.Obligations_Text
	Data_license.ModelComLimitations = license.Model.Com.Limitations_Text
	Data_license.ModelInternalRights = license.Model.Internal.Rights
	Data_license.ModelInternalObligations = license.Model.Internal.Obligations_Text
	Data_license.ModelInternalLimitations = license.Model.Internal.Limitations_Text
	Data_license.ModelOutputComRights = license.Model.Output_com.Rights
	Data_license.ModelOutputComObligations = license.Model.Output_com.Obligations_Text
	Data_license.ModelOutputComLimitations = license.Model.Output_com.Limitations_Text
	Data_license.ModelPublishRights = license.Model.Publish.Rights
	Data_license.ModelPublishObligations = license.Model.Publish.Obligations_Text
	Data_license.ModelPublishLimitations = license.Model.Publish.Limitations_Text
	Data_license.ModelResearchRights = license.Model.Research.Rights
	Data_license.ModelResearchObligations = license.Model.Research.Obligations_Text
	Data_license.ModelResearchLimitations = license.Model.Research.Limitations_Text
	Data_license.ModelRevRights = license.Model.Rev.Rights
	Data_license.ModelRevObligations = license.Model.Rev.Obligations_Text
	Data_license.ModelRevLimitations = license.Model.Rev.Limitations_Text

	if err := database.DB.Create(&Data_license).Error; err != nil {
		return nil, err
	}
	return
}

func GetDatalicensesByPage(p *utils.Pagination, t string) (Datalicenses []Datalicense, err error) {
	var total int64
	if t == "all" {
		err = database.DB.Model(&Datalicense{}).Scopes(p.GormPaginate()).Find(&Datalicenses).Error
		database.DB.Model(&Datalicense{}).Count(&total)
	} else {
		err = database.DB.Model(&Datalicense{}).Where("license_type = ?", t).Scopes(p.GormPaginate()).Find(&Datalicenses).Error
		database.DB.Model(&Datalicense{}).Where("license_type = ?", t).Count(&total)
	}
	if err != nil {
		return nil, err
	}
	p.Total = cast.ToInt(total)
	return
}

func GetDatalicenseBasicByID(id int) (Licensebasic *LicenseBasic, err error) {
	var _Datalicense Datalicense
	Licensebasic = new(LicenseBasic)
	err = database.DB.Model(&Datalicense{}).First(&_Datalicense, id).Error
	if err != nil {
		return nil, err
	}
	Licensebasic.Id = _Datalicense.Id
	Licensebasic.LicenseName = _Datalicense.LicenseName
	Licensebasic.LicenseType = _Datalicense.LicenseType
	Licensebasic.LicenseUuid = _Datalicense.LicenseUuid
	Licensebasic.OsiApproved = _Datalicense.OsiApproved
	Licensebasic.ShortIdentifier = _Datalicense.ShortIdentifier
	return
}

func GetDatalicenseBasicByName(name string) (Licensebasic *LicenseBasic, err error) {
	var _Datalicense Datalicense
	Licensebasic = new(LicenseBasic)
	err = database.DB.Model(&Datalicense{}).Where("license_name = ?", name).First(&_Datalicense).Error
	if err != nil {
		return nil, err
	}
	Licensebasic.Id = _Datalicense.Id
	Licensebasic.LicenseName = _Datalicense.LicenseName
	Licensebasic.LicenseType = _Datalicense.LicenseType
	Licensebasic.LicenseUuid = _Datalicense.LicenseUuid
	Licensebasic.OsiApproved = _Datalicense.OsiApproved
	Licensebasic.ShortIdentifier = _Datalicense.ShortIdentifier
	return
}

func SearchDatalicenseBasicByName(name string, t string) (LicenseBasics []LicenseBasic, err error) {
	if t == "all" {
		err = database.DB.Model(&Datalicense{}).Where("license_name LIKE ?", "%"+name+"%").Limit(20).Find(&LicenseBasics).Error
	} else {
		err = database.DB.Model(&Datalicense{}).Where("license_name LIKE ? AND license_type = ?", "%"+name+"%", t).Limit(20).Find(&LicenseBasics).Error
	}
	if err != nil {
		return nil, err
	}
	return
}

func GetDatalicenseDataByID(id int) (LicenseDataBoxs *LicenseData, err error) {
	var _Datalicense Datalicense
	var cans []LicenseDataCan
	var cannots []LicenseDataCannot
	var obligations []LicenseDataObligation
	var limitations []LicenseDataLimitation
	can := new(LicenseDataCan)
	cannot := new(LicenseDataCannot)
	obligation := new(LicenseDataObligation)
	limitation := new(LicenseDataLimitation)
	LicenseDataBoxs = new(LicenseData)
	_id_can, _id_cannot, _id_obligation, _id_limitation := 1, 1, 1, 1
	err = database.DB.Model(&Datalicense{}).First(&_Datalicense, id).Error
	if err != nil {
		return nil, err
	}
	o := reflect.ValueOf(_Datalicense)
	if o.Kind() == reflect.Ptr {
		o = o.Elem()
	}
	for i := 0; i < o.NumField(); i++ {
		_keyName := o.Type().Field(i).Name
		_keyValue := o.FieldByName(_keyName).String()
		if strings.Contains(_keyName, "Rights") && strings.Contains(_keyName, "Data") {
			_Property := strings.Split(_keyName, "Rights")
			if _keyValue != "No" {
				can.Id = _id_can
				can.Property = _Property[0]
				cans = append(cans, *can)
				_id_can++
			} else {
				cannot.Id = _id_cannot
				cannot.Property = _Property[0]
				cannots = append(cannots, *cannot)
				_id_cannot++
			}
		}
		if strings.Contains(_keyName, "Obligations") && strings.Contains(_keyName, "Data") {
			_Property := strings.Split(_keyName, "Obligations")
			if _keyValue != "No" {
				obligation.Id = _id_obligation
				obligation.Property = _Property[0]
				obligations = append(obligations, *obligation)
				_id_obligation++
			}
		}
		if strings.Contains(_keyName, "Limitations") && strings.Contains(_keyName, "Data") {
			_Property := strings.Split(_keyName, "Limitations")
			if _keyValue != "No" {
				limitation.Id = _id_limitation
				limitation.Property = _Property[0]
				limitations = append(limitations, *limitation)
				_id_limitation++
			}
		}
	}
	LicenseDataBoxs.Can = cans
	LicenseDataBoxs.Cannot = cannots
	LicenseDataBoxs.Obligation = obligations
	LicenseDataBoxs.Limitation = limitations

	return
}

func GetDatalicenseModelByID(id int) (LicenseModelBoxs *LicenseModel, err error) {
	var _Datalicense Datalicense
	var cans []LicenseModelCan
	var cannots []LicenseModelCannot
	var obligations []LicenseModelObligation
	var limitations []LicenseModelLimitation
	can := new(LicenseModelCan)
	cannot := new(LicenseModelCannot)
	obligation := new(LicenseModelObligation)
	limitation := new(LicenseModelLimitation)
	LicenseModelBoxs = new(LicenseModel)
	_id_can, _id_cannot, _id_obligation, _id_limitation := 1, 1, 1, 1
	err = database.DB.Model(&Datalicense{}).First(&_Datalicense, id).Error
	if err != nil {
		return nil, err
	}
	o := reflect.ValueOf(_Datalicense)
	if o.Kind() == reflect.Ptr {
		o = o.Elem()
	}
	for i := 0; i < o.NumField(); i++ {
		_keyName := o.Type().Field(i).Name
		_keyValue := o.FieldByName(_keyName).String()
		if strings.Contains(_keyName, "Rights") && strings.Contains(_keyName, "Model") {
			_Property := strings.Split(_keyName, "Rights")
			if _keyValue != "No" {
				can.Id = _id_can
				can.Property = _Property[0]
				cans = append(cans, *can)
				_id_can++
			} else {
				cannot.Id = _id_cannot
				cannot.Property = _Property[0]
				cannots = append(cannots, *cannot)
				_id_cannot++
			}
		}
		if strings.Contains(_keyName, "Obligations") && strings.Contains(_keyName, "Model") {
			_Property := strings.Split(_keyName, "Obligations")
			if _keyValue != "No" {
				obligation.Id = _id_obligation
				obligation.Property = _Property[0]
				obligations = append(obligations, *obligation)
				_id_obligation++
			}
		}
		if strings.Contains(_keyName, "Limitations") && strings.Contains(_keyName, "Model") {
			_Property := strings.Split(_keyName, "Limitations")
			if _keyValue != "No" {
				limitation.Id = _id_limitation
				limitation.Property = _Property[0]
				limitations = append(limitations, *limitation)
				_id_limitation++
			}
		}
	}
	LicenseModelBoxs.Can = cans
	LicenseModelBoxs.Cannot = cannots
	LicenseModelBoxs.Obligation = obligations
	LicenseModelBoxs.Limitation = limitations

	return
}

func GetDatalicenseOtherByID(id int) (LicenseOtherBoxs []LicenseOther, err error) {
	var _Datalicense Datalicense
	_id_other := 1
	other := new(LicenseOther)
	err = database.DB.Model(&Datalicense{}).First(&_Datalicense, id).Error
	if err != nil {
		return nil, err
	}
	o := reflect.ValueOf(_Datalicense)
	if o.Kind() == reflect.Ptr {
		o = o.Elem()
	}
	for i := 0; i < o.NumField(); i++ {
		_keyName := o.Type().Field(i).Name
		if strings.Contains(_keyName, "Other") {
			_Property := strings.Split(_keyName, "Other")
			other.Id = _id_other
			other.Property = _Property[1]
			LicenseOtherBoxs = append(LicenseOtherBoxs, *other)
			_id_other++
		}
	}
	return
}

func GetDatalicenseIndex(t string) (Licenseindexes []LicenseIndex, err error) {
	if t == "all" {
		err = database.DB.Model(&Datalicense{}).Find(&Licenseindexes).Error

	} else {
		err = database.DB.Model(&Datalicense{}).Where("license_type = ?", t).Find(&Licenseindexes).Error

	}
	if err != nil {
		return nil, err
	}
	return
}
