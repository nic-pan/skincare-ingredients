package models

import (
	"github.com/jinzhu/gorm"

	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

type SkinTypeIngredient struct {
	gorm.Model
	gorm.JoinTableHandler
	IngredientID uint
	SkinTypeID   uint
	Effect       string
}

func (SkinTypeIngredient) BeforeCreate(db *gorm.DB) error {
	db.SetJoinTableHandler(&Ingredient{}, "SkinType", &SkinTypeIngredient{})
	return nil
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&SkinTypeIngredient{})
}

func (SkinTypeIngredient *SkinTypeIngredient) CreateSkinTypeIngredient() *SkinTypeIngredient {
	db.NewRecord(SkinTypeIngredient)
	db.Create(&SkinTypeIngredient)
	return SkinTypeIngredient
}

func GetAllSkinTypeIngredients() []SkinTypeIngredient {
	var SkinTypeIngredients []SkinTypeIngredient
	db.Find(&SkinTypeIngredients)
	return SkinTypeIngredients
}

func GetSkinTypeIngredientById(Id int64) (*SkinTypeIngredient, *gorm.DB) {
	var SkinTypeIngredient SkinTypeIngredient
	db := db.Where("ID=?", Id).Find(&SkinTypeIngredient)
	return &SkinTypeIngredient, db
}

func DeleteSkinTypeIngredient(Id int64) *SkinTypeIngredient {
	var SkinTypeIngredient SkinTypeIngredient
	db.Where("ID=?", Id).Delete(SkinTypeIngredient)
	return &SkinTypeIngredient
}
