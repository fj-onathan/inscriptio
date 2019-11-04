package roles

import (
	"github.com/gin-gonic/gin"
	"inscriptio/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	{
		roles := r.Group("/roles")
		{
			roles.POST("/", middlewares.Authorized(1), create)
			roles.GET("/", middlewares.Authorized(1), list)
			roles.DELETE("/:role", middlewares.Authorized(1), remove)
			roles.PATCH("/:role", middlewares.Authorized(1), update)
		}
	}
}