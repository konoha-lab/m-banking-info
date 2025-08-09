package repository

import (
	"m-banking.info/db"
	"m-banking.info/model"
)

func CreateAccount(account *model.M_Account) error {
	return db.DB.Create(account).Error
}

func GetAccount() ([]model.M_Account, error) {
	var accounts []model.M_Account
	err := db.DB.Find(&accounts).Error

	return accounts, err
}
