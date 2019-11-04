package api

import (
	"github.com/gin-gonic/gin"
	_ "mailserver/server/api/logs"
	"mailserver/server/api/send"
	"reflect"
)

// Ping connection
func ping(c *gin.Context) {
	type Empty struct{}
	c.JSON(200, gin.H{
		"message": "do a pong on package: " + reflect.TypeOf(Empty{}).PkgPath(),
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	api := r.Group("/api")
	{
		// Helpers routing
		api.GET("/ping", ping)

		// Mail routing
		send.ApplyRoutes(api)
		//logs.ApplyRoutes(api)
	}
}