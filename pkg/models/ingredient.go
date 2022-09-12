package models

import (
	"github.com/jinzhu/gorm"
)

type Ingredient struct {
	gorm.Model
	Name      string     `gorm:"unique;not null;type:varchar(100)"json:name`
	SkinTypes []SkinType `gorm:"many2many:skintypes_ingredients"json:skinTypes`
	Effect    string     `json:effect`
	Slug      string     `json:slug`
}

func (Ingredient *Ingredient) CreateIngredient() (*Ingredient, error) {
	db.NewRecord(Ingredient)
	// str, _ := json.Marshal(Ingredient)
	// fmt.Printf("Creating %s", str)
	db := db.Create(&Ingredient)
	err := db.Error
	return Ingredient, err
}

func (Ingredient *Ingredient) UpdateIngredient() (*Ingredient, error) {
	err := db.Save(&Ingredient).Error
	return Ingredient, err
}

func GetAllIngredients() []Ingredient {
	var Ingredients []Ingredient
	db.Find(&Ingredients)
	return Ingredients
}

func GetIngredientById(Id uint) (*Ingredient, *gorm.DB) {
	var Ingredient Ingredient
	db := db.Where("ID=?", Id).Find(&Ingredient)
	return &Ingredient, db
}

func GetIngredientByName(Name string) (*Ingredient, *gorm.DB) {
	var Ingredient Ingredient
	db := db.Where("Name=?", Name).Find(&Ingredient)
	return &Ingredient, db
}

func DeleteIngredient(Id uint) *gorm.DB {
	var Ingredient Ingredient
	db := db.Where("ID=?", Id).Unscoped().Delete(&Ingredient)
	return db
}
