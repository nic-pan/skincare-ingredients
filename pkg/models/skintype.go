package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

var db *gorm.DB

type SkinType struct {
	gorm.Model
	ID   int    `gorm:""json:"id"`
	Name string `json:"name"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&SkinType{})
}

func (st *SkinType) CreateSkinType() *SkinType {
	db.NewRecord(st)
	db.Create(&st)
	return st
}

func GetAllSkinTypes() []SkinType {
	var SkinTypes []SkinType
	db.Find(&SkinTypes)
	return SkinTypes
}

func GetSkinTypeById(Id int64) (*SkinType, *gorm.DB) {
	var skinType SkinType
	db := db.Where("ID=?", Id).Find(&skinType)
	return &skinType, db
}

func DeleteSkinType(Id int64) *SkinType {
	var skinType SkinType
	db.Where("ID=?", Id).Delete(skinType)
	return &skinType
}
