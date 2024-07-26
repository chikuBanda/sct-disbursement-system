package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sct-system/models"
	"sct-system/repositories"
)

func registerNewBeneficiary(c *gin.Context) {
	var newBeneficiary models.Beneficiary

	if err := c.ShouldBindJSON(&newBeneficiary); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateBeneficiary(&newBeneficiary); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBeneficiary)
}

func getBeneficiaries(c *gin.Context) {
	beneficiaries, err := repositories.GetBeneficiaries()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, beneficiaries)
}
