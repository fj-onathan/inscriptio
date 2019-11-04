package models

import (
	"github.com/jinzhu/gorm"
	"inscriptio/libraries/common"
)

// User data model
type User struct {
	gorm.Model
	Username     string `sql:"type:varchar(45);"`
	DisplayName  string `sql:"type:varchar(100);"`
	PasswordHash string
	Email        string `sql:"type:varchar(150);"`
	RoleID       uint   `sql:"default: 1"`
	Role         Role   `json:"-"`
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"email":        u.Email,
		"display_name": u.DisplayName,
		"avatar":       string(u.DisplayName[0:2]),
		"role":         u.RoleID,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.DisplayName = m["display_name"].(string)
	u.RoleID = uint(m["role"].(float64))
}
