package models

import "gorm.io/gorm"

type Option func(*gorm.DB)

func ID(id int) Option {
	return func(db *gorm.DB) {
		if id != 0 {
			db.Where("`id` = ?", id)
		}
	}
}
func LicenseName(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`license_name` = ?", name)
		}
	}
}
func DataAccessRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`data_access_rights` = ?", name)
		}
	}
}
func DataTaggingRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`data_tagging_rights` = ?", name)
		}
	}
}
func DataRepresentRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`data_represent_rights` = ?", name)
		}
	}
}
func ModelBenchmarkRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`model_benchmark_rights` = ?", name)
		}
	}
}
func ModelResearchRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`model_research_rights` = ?", name)
		}
	}
}
func ModelPublishRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`model_publish_rights` = ?", name)
		}
	}
}
func ModelInternalRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`model_Internal_rights` = ?", name)
		}
	}
}
func ModelOutputComRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`model_output_com_rights` = ?", name)
		}
	}
}
func ModelComRights(name string) Option {
	return func(db *gorm.DB) {
		if name != "" {
			db.Where("`model_com_rights` = ?", name)
		}
	}
}
func ValidVarchar(key string, value string) Option {
	return func(db *gorm.DB) {
		if value != "" {
			db.Where("`"+key+"` = ?", value)
		}
	}
}
