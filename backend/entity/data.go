package entity

import (
	"gorm.io/gorm"
	//"github.com/waranyafeen/project/backend/entity"
)


func SetupData(db *gorm.DB) {
	var role []Role
	db.Find(&role, 100)
	if len(role) > 0 {
		return
	}

	// roles
	roles := []*Role{
		{
			gormModel: gormModel{ID: 100},
			Name:      "admin",
		},
		{
			gormModel: gormModel{ID: 101},
			Name:      "user",
		},
	}
	db.Create(roles)

	// employee
	employees := []Employee{
		{
			gormModel: gormModel{ID: 201},
			Firstname:  "manager",
			Lastname:   "manager",
			Email:      "manager@manager.com",
			Password:   "manager",
			Tel:      	"0988888888",
			PositionID: 201,
			PrecedeID:  1,
			GenderID:   1,
		},
		{
			gormModel: gormModel{ID: 202},
			Firstname:  "admin",
			Lastname:   "admin",
			Email:      "admin@admin.com",
			Password:   "admin",
			Tel:      	"0988888888",
			PositionID: 202,
			PrecedeID:  2,
			GenderID:   2,
		},	
	}
	db.Create(&employees)


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

	// gender data
	genders := []Gender{
		{
			gormModel: gormModel{ID: 1},
			Name:      "ชาย",
		},
		{
			gormModel: gormModel{ID: 2},
			Name:      "หญิง",
		},
	}

	db.Create(&genders)

	// precedes data
	precedes := []Precede{
		{
			gormModel: gormModel{ID: 1},
			Name:      "นาย",
		},
		{
			gormModel: gormModel{ID: 2},
			Name:      "นาง",
		},
		{
			gormModel: gormModel{ID: 3},
			Name:      "นางสาว",
		},
	}
	db.Create(&precedes)


	//car
	cars := []Car{
		{
			gormModel: gormModel{ID: 1},
			Name:      "ป.1/1 ก.",
			Route: "นครราชสีมา - ขอนแก่น",
		},
		{
			gormModel: gormModel{ID: 2},
			Name:      "ป.1/2 ก.",
			Route: "นครราชสีมา - สระบุรี",
		},
		{
			gormModel: gormModel{ID: 3},
			Name:      "ป.1/3 ก.",
			Route: "นครราชสีมา - กรุงเทพฯ",
		},
		{
			gormModel: gormModel{ID: 4},
			Name:      "ป.1/4 ก.",
			Route: "นครราชสีมา - บุรีรัมย์",
		},
	}
	db.Create(cars)

	//ticket
	tickets := []Ticket{
		{
			gormModel: gormModel{ID: 1},
			Price:      150,
			CarID: 1,
		},
		{
			gormModel: gormModel{ID: 2},
			Price:      250,
			CarID: 2,
		},
		{
			gormModel: gormModel{ID: 3},
			Price:      375,
			CarID: 3,
		},
		{
			gormModel: gormModel{ID: 4},
			Price:      95,
			CarID: 4,
		},
	}
	db.Create(&tickets)
}