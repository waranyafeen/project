package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/project/backend/entity"
	"github.com/asaskevich/govalidator"
)

type EmployeeForUpdate struct {
	
}

// GET /employees
func GetAllEmployees(c *gin.Context) {
	var employees []entity.Employee

	if err := entity.DB().Joins("Gender").Joins("Position").Joins("Precede").Find(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})

}

// GET /employees/id
func GetEmployees(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")

	if err := entity.DB().Joins("Gender").Joins("Position").Joins("Precede").First(&employee, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})

}

// POST /employees
func CreateEmployees(c *gin.Context) {
	var employee entity.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Create employee successfully"})

}

// PUT /employee/id
func UpdateEmployees(c *gin.Context) {
	var employee EmployeeForUpdate
	id := c.Param("id")

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Table("employee").Where("id = ?", id).Updates(employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Updated your employee successfully"})

}

// DELETE /employees/:id
func DeleteEmployees(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")

	if rows := entity.DB().Delete(&employee, id).RowsAffected; rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Delete your employee successfully"})

}

// GET /employee/position
func GetAllPositions(c *gin.Context){ 
	var positions []entity.Position

	if err := entity.DB().Find(&positions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": positions})
}

// GET /employee/precede
func GetAllPrecedes(c *gin.Context){ 
	var precedes []entity.Precede

	if err := entity.DB().Find(&precedes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": precedes})
}

// GET /employee/gender
func GetAllGendersForEmployee(c *gin.Context) {
	var genders []entity.Gender

	if err := entity.DB().Find(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genders})
}