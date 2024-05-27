package models

import "time"

type Journal struct {
	Id        int       `gorm:"primaryKey;"`
	Date      time.Time `json:"date"`
	IsDeposit bool      `json:"is_deposit"`
	Amount    float64   `json:"amount"`

	AccountNumber string `json:"account_number"`
}
