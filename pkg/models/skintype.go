package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

var db *gorm.DB

type SkinType struct {
	gorm.Model
	Name            string `gorm:"unique;not null;type:varchar(100);default:null"json:"name"`
	Characteristics string `json:"characteristics"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&SkinType{})
}

func (st *SkinType) CreateSkinType() (*SkinType, error) {
	db.NewRecord(st)
	err := db.Create(&st).Error
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

func GetSkinTypeById(Id int64) (*SkinType, *gorm.DB) {
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

func DeleteSkinType(Id int64) *SkinType {
	var skinType SkinType
	db.Where("ID=?", Id).Delete(skinType)
	return &skinType
}
