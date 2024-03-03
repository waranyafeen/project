package controllers

import(
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/project/backend/entity"
	"github.com/asaskevich/govalidator"
)

// GET /payments
func GetAllPayments(c *gin.Context) {
	var payments []entity.Payment

	if err := entity.DB().Joins("Ticket").Joins("User").Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})

}

// GET /payments/id
func GetPayments(c *gin.Context) {
	var payment entity.Payment
	id := c.Param("id")

	if err := entity.DB().Joins("Ticket").Joins("User").First(&payment, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})

}

// POST /payments
func CreatePayments(c *gin.Context) {
	var payment entity.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Create payment successfully"})

}

// PUT /payment/id
func UpdatePayments(c *gin.Context) {
	var payment entity.Payment
	id := c.Param("id")

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment.ID = uint(idUint)

	if err := entity.DB().Save(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated your payment successfully"})

}

// DELETE /payments/id
func DeletePayments(c *gin.Context) {
	var payment entity.Payment
	id := c.Param("id")

	if rows := entity.DB().Delete(&payment, id).RowsAffected; rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Delete your payment successfully"})

}