package models

import (
	"github.com/nic_pan/skincare-ingredients/pkg/config"
)

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Combination{})
}
