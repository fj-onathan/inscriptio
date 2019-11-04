package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Role{},
		&Company{},
		&Log{},
	)
	// set up foreign keys
	db.Model(&User{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")
	db.Model(&Company{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Log{}).AddForeignKey("email_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}
