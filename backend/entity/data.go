package entity

import (

	"gorm.io/gorm"
)

func SetupData(db *gorm.DB) {
	var role []Role
	db.Find(&role, 100)
	if len(role) > 0 {
		return
	}
}