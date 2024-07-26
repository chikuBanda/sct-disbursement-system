package router

import (
	"github.com/gin-gonic/gin"
)

func StartRouting() {
	router := gin.Default()

	router.POST("/register_beneficiary", registerNewBeneficiary)
	router.GET("/get_beneficiaries", getBeneficiaries)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
