package model

import (
	"go-hello/common"
	"go-hello/storage"
	"gorm.io/gorm"
)

type CV struct {
	gorm.Model
	UserId     uint   `gorm:"not null;" json:"user_id"`
	Name       string `gorm:"not null;" json:"name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Github     string `json:"github"`
	Linkedin   string `json:"linkedin"`
	Summary    string `json:"summary"`
	Skills     string `json:"skills"`
	Experience []*CVExperience
	Education  []*CVEducation
}

func (CV) TableName() string {
	return "cv"
}

type CVExperience struct {
	gorm.Model
	CVId           uint   `gorm:"not null;" json:"cv_id"`
	JobName        string `gorm:"not null;" json:"job_name"`
	Company        string `json:"company"`
	Period         string `json:"period"`
	JobDescription string `json:"job_description"`
}

func (CVExperience) TableName() string {
	return "cv_experience"
}

type CVEducation struct {
	gorm.Model
	CVId    uint    `gorm:"not null;" json:"cv_id"`
	name    string  `gorm:"not null;" json:"name"`
	faculty string  `json:"faculty"`
	gpa     float32 `json:"gpa"`
}

func (CVEducation) TableName() string {
	return "cv_education"
}

func ListCVByCondition(
	conditions map[string]interface{},
	filter *CVFilter,
	paging *common.Paging) ([]CV, error) {
	var result []CV
	db := storage.Database
	db = db.Table(CV{}.TableName()).Where(conditions)
	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("name ilike ?", v.Name)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (cv *CV) Create() (*CV, error) {
	err := storage.Database.Create(&cv).Error
	if err != nil {
		return &CV{}, err
	}
	return cv, nil
}
