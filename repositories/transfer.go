package repositories

import (
	"sct-system/database"
	"sct-system/models"
)

func GetTransfers() (error, []models.Transfer) {
	var transfers []models.Transfer
	err := database.DB.Model(&models.Transfer{}).Find(&transfers).Error
	return err, transfers
}

func GetTransfersByBeneficiary(beneficiary models.Beneficiary) (error, []models.Transfer) {
	var transfers []models.Transfer
	err := database.DB.Model(&models.Transfer{}).Where("beneficiary_id = ?", beneficiary.ID).Find(&transfers).Error
	return err, transfers
}

func CreateTransfer(beneficiaryId uint, newTransfer *models.Transfer) error {
	var beneficiary models.Beneficiary
	database.DB.Model(&models.Beneficiary{}).Where("id = ?", beneficiaryId).First(&beneficiary)

	err := database.DB.Model(&models.Beneficiary{}).Association("Transfers").Append(newTransfer)
	return err
}
