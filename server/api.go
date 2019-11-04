package api

import (
	"github.com/gin-gonic/gin"
	"inscriptio/server/api"
	"inscriptio/server/setup"
	"net/http"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	// index documentation
	r.LoadHTMLGlob("html/*/*")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "api")
	})
	r.Static("/assets", "./html/assets")
	r.GET("api", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "inscriptio documentation",
			"description": "GO email server to send transacional emails, api for developing.",
			"github": gin.H{
				"name": "Github",
				"link": "https://github.com/dev-fjonathan/inscriptio",
			},
			"documentation": gin.H{
				"name": "Documentation",
				"link": "https://github.com/dev-fjonathan/inscriptio/tree/master/documentation",
			},
		})
	})
	// routing by grouping
	server := r.Group("/")
	{
		api.ApplyRoutes(server)
		setup.ApplyRoutes(server)
	}
}
