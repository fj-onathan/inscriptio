package companies

import (
	"github.com/gin-gonic/gin"
	"mailserver/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(c *gin.RouterGroup) {
	{
		companies := c.Group("/companies")
		{
			companies.POST("/", middlewares.Authorized(2), create)
			companies.GET("/", middlewares.Authorized(2), list)
			companies.DELETE("/:company", middlewares.Authorized(2), remove)
			companies.PATCH("/:company", middlewares.Authorized(2), update)
		}
	}
}