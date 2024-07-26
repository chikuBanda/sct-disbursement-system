package repositories

import (
	"sct-system/database"
	"sct-system/models"
)

func GetBeneficiaries() ([]models.Beneficiary, error) {
	var beneficiaries []models.Beneficiary
	err := database.DB.Model(&models.Beneficiary{}).Find(&beneficiaries).Error
	return beneficiaries, err
}

func CreateBeneficiary(newBeneficiary *models.Beneficiary) error {
	err := database.DB.Create(&newBeneficiary).Error

	return err
}
