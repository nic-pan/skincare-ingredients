package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

type Combination struct {
	gorm.Model
	Ingredient1_id uint
	Ingredient2_id uint
	Ingredient1    Ingredient `gorm:"foreignKey:Ingredient1_id"json:ingredient1`
	Ingredient2    Ingredient `gorm:"foreignKey:Ingredient2_id"json:ingredient2`
	IsHealthy      bool       `json:isHealthy`
	Reason         string     `json:reason`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Combination{})
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

func GetCombinationById(Id int64) (*Combination, *gorm.DB) {
	var Combination Combination
	db := db.Where("ID=?", Id).Find(&Combination)
	return &Combination, db
}

func DeleteCombination(Id int64) *Combination {
	var Combination Combination
	db.Where("ID=?", Id).Delete(Combination)
	return &Combination
}
