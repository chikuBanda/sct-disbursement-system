package repositories

import (
	"gorm.io/gorm"
	"sct-system/database"
	"sct-system/models"
	"time"
)

func GetTransfers() ([]models.Transfer, error) {
	var transfers []models.Transfer
	err := database.DB.Model(&models.Transfer{}).Find(&transfers).Error
	return transfers, err
}

func GetTransfersByBeneficiary(beneficiary models.Beneficiary) (error, []models.Transfer) {
	var transfers []models.Transfer
	err := database.DB.Model(&models.Transfer{}).Where("beneficiary_id = ?", beneficiary.ID).Find(&transfers).Error
	return err, transfers
}

func CreateTransfer(newTransfer *models.Transfer) error {
	// Use an atomic transaction to create both the transfer and transaction history
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var beneficiary models.Beneficiary
		tx.Model(&models.Beneficiary{}).Where("id = ?", newTransfer.BeneficiaryId).First(&beneficiary)

		if err := tx.Model(&beneficiary).Association("Transfers").Append(newTransfer); err != nil {
			return err
		}

		var newTransactionHistory models.TransactionHistory
		newTransactionHistory.TransferID = newTransfer.ID
		newTransactionHistory.Timestamp = time.Now()
		newTransactionHistory.Status = "PAID"

		err := tx.Model(&models.TransactionHistory{}).Create(&newTransactionHistory).Error
		if err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return err
}
