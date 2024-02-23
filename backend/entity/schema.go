package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Firstname string	`gorm:"default:Employee" `
	Lastname  string	`gorm:"default:Employee" `
	Tel       string	`valid:"required~Tel is required,stringlength(10|10)~Tel must be at 10 characters"`
	Email     string	`gorm:"unique" valid:"required~Email is required,email~Invalid email"`
	//Username  string
	Password  string	`valid:"required~Password is required,minstringlength(4)~Password must be at 4 characters"`

	PositionID uint
	Position   *Position

	PrecedeID uint
	Precede   *Precede

	GenderID uint
	Gender   *Gender
}

type Precede struct {
	gorm.Model
	Name string `gorm:"unique"`

	Employee []Employee
}

type Position struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Salary      string

	Employee []Employee
}

type Gender struct {
	gorm.Model
	Name string	`gorm:"unique"`

	Employee []Employee
	User     []User
}

type User struct {
	gorm.Model
	Firstname string	`gorm:"default:User" `
	Lastname  string	`gorm:"default:User" `
	Age       int		`valid:"required~Age is required,gte=0~Age must be at least 0"`
	Phone     string	`valid:"required~Phone number is required,stringlength(10|10)~Phone must be at 10 characters"`
	Email     string	`gorm:"unique" valid:"required~Email is required,email~Invalid email"`
	//Username  string
	Password  string	`valid:"required~Password is required,minstringlength(4)~Password must be at 4 characters"`

	GenderID uint
	Gender   *Gender

	RoleID uint	
	Role   *Role

	Payment []Payment
}

type Role struct {
	gorm.Model
	Name string `gorm:"unique"`

	User []User
}

type Ticket struct {
	gorm.Model
	Price  string
	Seat   int	
	Detail string	

	CarID uint
	Car   *Car

	DepartureID uint
	Departure   *Departure

	Payment []Payment
}

type Departure struct {
	gorm.Model
	DepsrtureStation  string
	DestinationStaion string
	Date              time.Time

	Ticket []Ticket
}

type Car struct {
	gorm.Model
	Name     string	
	CarModel string
	Route    string

	Ticket []Ticket
}

type Payment struct {
	gorm.Model
	Total int

	TicketID uint
	Ticket   *Ticket

	UserID uint
	User   *User
}
