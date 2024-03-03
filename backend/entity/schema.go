package entity

import (
	"time"

	"gorm.io/gorm"
	//"github.com/asaskevich/govalidator"
)

type gormModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Employee struct {
	gormModel
	Firstname string `gorm:"default:Employee" `
	Lastname  string `gorm:"default:Employee" `
	Tel       string `valid:"required~Tel is required,stringlength(10|10)~Tel must be at 10 characters"`
	Email     string `gorm:"unique" valid:"required~Email is required,email~Invalid email"`
	Password  string `valid:"required~Password is required,minstringlength(4)~Password must be at 4 characters"`

	PositionID uint	` valid:"required~Position is required,refer=positions~Position does not exist"`
	Position   *Position

	PrecedeID uint
	Precede   *Precede

	GenderID uint
	Gender   *Gender
}

type Precede struct {
	gormModel
	Name string `gorm:"unique"`

	Employee []Employee
}

type Position struct {
	gormModel
	Name   string `gorm:"unique"`
	Salary int

	Employee []Employee
}

type Gender struct {
	gormModel
	Name string `gorm:"unique"`

	Employee []Employee
	User     []User
}

type User struct {
	gormModel
	Firstname string `gorm:"default:User" `
	Lastname  string `gorm:"default:User" `
	Age       int    `valid:"required~Age is required,gte=0~Age must be at least 0"`
	Phone     string `valid:"required~Phone number is required,stringlength(10|10)~Phone must be at 10 characters"`
	Email     string `gorm:"unique" valid:"required~Email is required,email~Invalid email"`
	Password  string `valid:"required~Password is required,minstringlength(4)~Password must be at 4 characters"`

	GenderID uint
	Gender   *Gender

	RoleID uint `gorm:"default:101"`
	Role   *Role

	Payment []Payment
}

type Role struct {
	gormModel
	Name string `gorm:"unique"`

	User     []User
}

type Ticket struct {
	gormModel
	Price  int
	Seat   int
	Detail string

	CarID uint
	Car   *Car

	DepartureID uint
	Departure   *Departure

	Payment []Payment
}

type Province struct {
	gormModel
	Name string `gorm:"unique"`
}

type Departure struct {
	gormModel
	DepsrtureStation  string
	DestinationStaion string
	Date              time.Time

	Ticket []Ticket
}

type Car struct {
	gormModel
	Name  string
	Route string //เส้นทาง

	Ticket []Ticket
}

type Payment struct {
	gormModel
	Total int

	TicketID uint
	Ticket   *Ticket

	UserID uint
	User   *User
}
