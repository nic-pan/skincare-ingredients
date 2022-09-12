package models

import (
	"github.com/jinzhu/gorm"
)

type Combination struct {
	gorm.Model
	Ingredient1  uint   `gorm:"foreignKey:Ingredient1_id"json:ingredient1_id`
	Ingredient2  uint   `gorm:"foreignKey:Ingredient2_id"json:ingredient2_id`
	IsCompatible bool   `json:isCompatible`
	Reason       string `json:reason`
}

func (Combination *Combination) CreateCombination() *Combination {
	db.NewRecord(Combination)
	db.Create(&Combination)
	return Combination
}

func GetAllCombinations() []Combination {
	var Combinations []Combination
	db.Find(&Combinations)
	return Combinations
}

func GetCombinationById(Id uint) (*Combination, *gorm.DB) {
	var Combination Combination
	db := db.Where("ID=?", Id).Find(&Combination)
	return &Combination, db
}

func GetCombinationForIngredients(ingr1 *Ingredient, ingr2 *Ingredient) *Combination {
	var Combination Combination
	db.Where("Ingredient1.ID=? AND Ingredient2.ID=?", ingr1.ID, ingr2.ID).Find(&Combination)
	return &Combination
}

func DeleteCombination(Id uint) *Combination {
	var Combination Combination
	db.Where("ID=?", Id).Delete(Combination)
	return &Combination
}
