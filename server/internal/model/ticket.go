package model

type Ticket struct {
	ID   uint   `gorm:"primaryKey;default:auto_random()"`
	Code string `gorm:"unique;not null"`
	User string
}
