package models

import (
	"github.com/jinzhu/gorm"
	"inscriptio/libraries/common"
)

// Log data model
type Log struct {
	gorm.Model
	IP	     string `sql:"type:varchar(50);"`
	Message  string `sql:"type:text;"`
	Email    User `gorm:"foreignkey:UserID"`
	EmailID  uint
}

// Serialize serializes log data
func (l *Log) Serialize() common.JSON {
	return common.JSON{
		"id":    l.ID,
		"ip":  	 l.IP,
		"message": l.Message,
		"email":  l.EmailID,
	}
}

func (l *Log) Read(m common.JSON) {
	l.ID = uint(m["id"].(float64))
	l.IP = m["ip"].(string)
	l.Message = m["message"].(string)
	l.EmailID = m["email"].(uint)
}
