package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

type Ingredient struct {
	gorm.Model
	Name         string     `gorm:"unique"json:name`
	ForSkinTypes []SkinType `json:skinTypes`
	Effect       string     `json:effect`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Ingredient{})
}

func (Ingredient *Ingredient) CreateIngredient() *Ingredient {
	db.NewRecord(Ingredient)
	db.Create(&Ingredient)
	return Ingredient
}

func GetAllIngredients() []Ingredient {
	var Ingredients []Ingredient
	db.Find(&Ingredients)
	return Ingredients
}

func GetIngredientById(Id int64) (*Ingredient, *gorm.DB) {
	var Ingredient Ingredient
	db := db.Where("ID=?", Id).Find(&Ingredient)
	return &Ingredient, db
}

func DeleteIngredient(Id int64) *Ingredient {
	var Ingredient Ingredient
	db.Where("ID=?", Id).Delete(Ingredient)
	return &Ingredient
}
