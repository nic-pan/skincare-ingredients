package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type SkinType struct {
	gorm.Model
	Name            string `gorm:"unique;not null;type:varchar(100)"json:"name"`
	Characteristics string `json:"characteristics"`
	Slug            string `json:slug`
}

func (st *SkinType) CreateSkinType() (*SkinType, error) {
	db.NewRecord(st)
	db := db.Create(&st)
	err := db.Error
	return st, err
}

func (st *SkinType) UpdateSkinType() (*SkinType, error) {
	err := db.Save(&st).Error
	return st, err
}

func GetAllSkinTypes() []SkinType {
	var SkinTypes []SkinType
	db.Find(&SkinTypes)
	return SkinTypes
}

func GetSkinTypeById(Id uint) (*SkinType, *gorm.DB) {
	var skinType SkinType
	db := db.Where("ID=?", Id).First(&skinType)
	if err := db.Error; err != nil && err.Error() == "record not found" {
		return nil, db
	}
	return &skinType, db
}

func GetSkinTypeByName(Name string) (*SkinType, *gorm.DB) {
	var skinType SkinType
	db := db.Where("Name=?", Name).Find(&skinType)
	return &skinType, db
}

func DeleteSkinType(Id uint) *gorm.DB {
	var skinType SkinType = SkinType{}
	db := db.Where("ID=?", Id).Unscoped().Delete(&skinType)
	return db
}
