package models

import "time"

type Account struct {
	ID       int    `gorm:"primary_key"`
	Email    string `gorm:"unique"`
	Password string
}

type Contact struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Phone       string
	Description string
}

type Partner struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Description string
	Contacts    []Contact `gorm:"many2many:partner_contacts"`
}

type Bid struct {
	ID          int `gorm:"primary_key"`
	Description string
	Amount      int
	CreatedAt   time.Time
}
