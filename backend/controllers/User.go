package controllers

import(
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/project/backend/entity"
	"github.com/asaskevich/govalidator"
)

// GET /users
func GetAllUsers(c *gin.Context) {
	var users []entity.User

	if err := entity.DB().Where("role_id = ?", 101).Joins("Gender").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})

}

// GET /users/id
func GetUsers(c *gin.Context) {
	var user entity.User
	id := c.Param("id")

	if err := entity.DB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})

}

// POST /users
func CreateUsers(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Create user successfully"})

}

// PUT /user/id
func UpdateUsers(c *gin.Context) {
	var user entity.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(idUint)

	if err := entity.DB().Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated your user successfully"})

}

// DELETE /users/id
func DeleteUsers(c *gin.Context) {
	var user entity.User
	id := c.Param("id")

	if rows := entity.DB().Delete(&user, id).RowsAffected; rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Delete your user successfully"})

}

// GET /users/role
func GetAllRoles(c *gin.Context) {
	var roles []entity.Role

	if err := entity.DB().Find(&roles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// GET /users/gender
func GetAllGendersForUser(c *gin.Context) {
	var genders []entity.Gender

	if err := entity.DB().Find(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genders})
}