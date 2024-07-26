package repositories

import (
	"sct-system/database"
	"sct-system/models"
)

func GetTransactionHistoryList() (error, []models.TransactionHistory) {
	var transactions []models.TransactionHistory
	err := database.DB.Model(&models.TransactionHistory{}).Find(&transactions).Error
	return err, transactions
}

func GetTransactionHistoryById(id uint) (error, models.TransactionHistory) {
	var transaction models.TransactionHistory
	err := database.DB.Model(&models.TransactionHistory{}).Where("id = ?", id).First(&transaction).Error
	return err, transaction
}

func CreateTransactionHistory(transferId uint, newTransactionHistory *models.TransactionHistory) error {
	var transfer models.Transfer
	database.DB.Model(&models.Transfer{}).Where("id = ?", transferId).First(&transfer)

	err := database.DB.Model(&transfer).Association("TransactionHistories").Append(newTransactionHistory)
	return err
}
