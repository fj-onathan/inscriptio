package roles

import (
	"api-fuology/database/models/system"
	"api-fuology/libraries/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Role type alias
type Role = models.Role

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var roles []Role
	var users []User
	if err := db.Limit(10).Order("id asc").Find(&roles).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Not founded roles in system database.",
			},
		)
		return
	}
	length := len(roles)
	serialized := make([]JSON, length, length)
	for i := 0; i < length; i++ {
		db.Find(&User{}).Where("role_id = ?", roles[i].ID).Find(&users)
		roles[i].Statistics = uint(len(users))
		serialized[i] = roles[i].Read()
	}
	c.JSON(http.StatusOK,
		common.JSON{
			"success": true,
			"data":    serialized,
		},
	)
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name string `json:"name" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "Required inputs cannot be empty",
		})
		return
	}
	role := Role{Name: requestBody.Name}
	db.NewRecord(role)
	db.Create(&role)
	c.JSON(http.StatusOK,
		common.JSON{
			"success": true,
			"data":    role.Serialize(),
		},
	)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("role")
	type RequestBody struct {
		Name string `json:"name" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "Required inputs cannot be empty.",
			},
		)
		return
	}
	var role Role
	db.Where("id = ?", id).First(&role)
	role.Name = requestBody.Name
	db.Save(&role)
	c.JSON(http.StatusOK, common.JSON{
		"data":    role.Serialize(),
		"success": true,
	})
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("role")
	db.Where("id = ?", id).Delete(&Role{})
	db.Where("role_id = ?", id).Delete(&User{})
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
	})
}
