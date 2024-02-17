package model

type Ticket struct {
	ID   uint `gorm:"primaryKey;default:auto_random()"`
	Code string
	User string
}
