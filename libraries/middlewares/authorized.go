package middlewares

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mailserver/database/models"
	"net/http"
)

// Authorized master to do some requests
func Authorized(utype int) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		key := c.Query("key")
		requiredRole := utype
		var user models.User

		err := db.Preload("Role").Where("password_hash = ?", key).First(&user).Error
		if (err != nil) || (len(key) < 0) {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"error":   "Key code is wrong or not defined.",
				},
			)
			return
		}
		if int(user.RoleID) > requiredRole {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"error":   "Don't have permissions.",
				},
			)
			return
		}
	}
}
