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

	//province
	provinces := []Province {
		{
			gormModel: gormModel{ID: 1},
			Name:      "นครราชสีมา",	
			
		},
		{
			gormModel: gormModel{ID: 2},
			Name:      "ขอนแก่น",
		},
		{
			gormModel: gormModel{ID: 3},
			Name:      "กรุงเทพฯ",
		},
		{
			gormModel: gormModel{ID: 4},
			Name:      "บุรีรัมย์",
		},
		{
			gormModel: gormModel{ID: 5},
			Name:      "สระบุรี",
		},
	}
	db.Create(&provinces)
}