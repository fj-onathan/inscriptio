package send

import (
	"github.com/gin-gonic/gin"
	"mailserver/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	mails := r.Group("/send")
	{
		mails.POST("/normal", middlewares.Authorized(2), normal)
		//mails.GET("/", list)
	}
}
