package controllers

import(
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/project/backend/entity"
	"github.com/asaskevich/govalidator"
)

// GET /tickets
func GetAllTickets(c *gin.Context) {
	var tickets []entity.Ticket

	if err := entity.DB().Joins("Car").Joins("Departure").Find(&tickets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tickets})

}

// GET /tickets/id
func GetTickets(c *gin.Context) {
	var ticket entity.Ticket
	id := c.Param("id")

	if err := entity.DB().Joins("Car").Joins("Departure").First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ticket})

}

// POST /tickets
func CreateTickets(c *gin.Context) {
	var ticket entity.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Create ticket successfully"})

}

// PUT /ticket/id
func UpdateTickets(c *gin.Context) {
	var ticket entity.Ticket
	id := c.Param("id")

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket.ID = uint(idUint)

	if err := entity.DB().Save(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated your ticket successfully"})

}

// DELETE /tickets/id
func DeleteTickets(c *gin.Context) {
	var ticket entity.Ticket
	id := c.Param("id")

	if rows := entity.DB().Delete(&ticket, id).RowsAffected; rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Delete your ticket successfully"})

}

// GET /tickets/provice
func GetAllProvinces(c *gin.Context) {
	var provinces []entity.Province

	if err := entity.DB().Find(&provinces).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": provinces})
}


// GET /tickets/car
func GetAllCars(c *gin.Context) {
	var cars []entity.Car

	if err := entity.DB().Find(&cars).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// GET /ticket/departure
func GetAllDepartures(c *gin.Context) {
	var departures []entity.Departure

	if err := entity.DB().Find(&departures).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": departures})
}