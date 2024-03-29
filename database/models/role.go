package models

import (
	"github.com/jinzhu/gorm"
	"inscriptio/libraries/common"
)

// Role data model
type Role struct {
	gorm.Model
	Name       string
	Statistics uint `sql:"-"`
}

// Serialize serializes role data
func (u *Role) Serialize() common.JSON {
	return common.JSON{
		"id":   u.ID,
		"name": u.Name,
	}
}

// Read role data
func (p *Role) Read() common.JSON {
	return common.JSON {
		"id":    p.ID,
		"name":  p.Name,
		"users": p.Statistics,
	}
}
