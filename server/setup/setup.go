package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/t-tiger/gorm-bulk-insert"
	"inscriptio/database/models"
	"inscriptio/server/setup/auth"
	"inscriptio/server/setup/companies"
	"inscriptio/server/setup/roles"
	"os"
	"reflect"
)

// Ping connection
func ping(c *gin.Context) {
	type Empty struct{}
	c.JSON(200, gin.H{
		"message": "do a pong on package: " + reflect.TypeOf(Empty{}).PkgPath(),
	})
}

// Install
func install(c *gin.Context) {
	// GORM: where is bulk insert?
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("mysql", dbConfig)

	var rolesRecords []interface{}
	rolesRecords = append(rolesRecords, models.Role{Name: "Master"})
	rolesRecords = append(rolesRecords, models.Role{Name: "Admin"})

	err = gormbulk.BulkInsert(db, rolesRecords, 10)
	if err != nil {
		// do something
	}

	c.JSON(200, gin.H{
		"message": "Dependencies installed succesfully",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	setup := r.Group("/setup")
	{
		// Helpers routing
		setup.GET("/ping", ping)
		setup.GET("/install", install)

		// Roles routing
		roles.ApplyRoutes(setup)
		auth.ApplyRoutes(setup)
		companies.ApplyRoutes(setup)
	}
}