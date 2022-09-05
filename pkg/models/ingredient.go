package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

type Ingredient struct {
	gorm.Model
	Name      string     `gorm:"unique;not null;type:varchar(100)"json:name`
	SkinTypes []SkinType `gorm:"many2many:skintypes_ingredients"json:skinTypes`
	Effect    string     `json:effect`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Ingredient{})
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

func DeleteIngredient(Id uint) *gorm.DB {
	var Ingredient Ingredient
	db := db.Where("ID=?", Id).Unscoped().Delete(&Ingredient)
	return db
}
