package model

import "gorm.io/gorm"

type M_Account struct {
	gorm.Model

	AccountNumber     int64   `gorm:"not null:unique" json:"account_number"`
	TotalBalance      float64 `gorm:"type:decimal(15,2)" json:"total_balance"`
	AccountStatus     int     `json:"account_status"`
	AccountStatusDesc string  `json:"account_status_desc"`
	BlockStatus       bool    `json:"block_status"`
}
