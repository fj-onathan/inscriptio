package companies

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mailserver/database/models"
	"mailserver/libraries/common"
	"net/http"
)

// Company alias
type Company = models.Company

// JSON type alias
type JSON = common.JSON

func list(c * gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	var companies []Company
	if err := db.Limit(10).Order("id asc").Find(&companies).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Not founded companies in system database.",
			},
		)
		return
	}
	length := len(companies)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = companies[i].Serialize()
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
		Email string `json:"email" binding:"required"`
		Logo string `json:"logo"`
		Color string `json:"color"`
		Code string `json:"code" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H {
			"error": "Required inputs cannot be empty",
		})
		return
	}
	company := Company{
		Name: requestBody.Name,
		Email: requestBody.Email,
		Logo: requestBody.Logo,
		Color: requestBody.Color,
		Code: requestBody.Code,
	}
	db.NewRecord(company)
	db.Create(&company)
	c.JSON(http.StatusOK,
		common.JSON{
			"success": true,
			"data":    company.Serialize(),
		},
	)
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("company")
	db.Where("code = ?", id).Delete(&Company{})
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
	})
}


func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("company")
	type RequestBody struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Logo string `json:"logo"`
		Color string `json:"color"`
		Code string `json:"code" binding:"required"`
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
	var company Company
	db.Where("code = ?", id).First(&company)
	company.Name = requestBody.Name
	company.Email = requestBody.Email
	company.Logo = requestBody.Logo
	company.Color = requestBody.Color
	company.Code = requestBody.Code
	db.Save(&company)
	c.JSON(http.StatusOK, common.JSON{
		"data":    company.Serialize(),
		"success": true,
	})
}
