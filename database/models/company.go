package models

import (
	"github.com/jinzhu/gorm"
	"inscriptio/libraries/common"
)

// Company data model
type Company struct {
	gorm.Model
	Name    string `sql:"type:varchar(100);"`
	Email   string `sql:"type:varchar(150);"`
	Logo    string
	Color   string `sql:"type:varchar(10);"`
	Code	string
	Allowed int
}

// Serialize serializes company data
func (u *Company) Serialize() common.JSON {
	return common.JSON{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"logo":  u.Logo,
		"color": u.Color,
		"code":  u.Code,
	}
}

func (u *Company) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Name = m["name"].(string)
	u.Email = m["email"].(string)
	u.Logo = m["logo"].(string)
	u.Color = m["color"].(string)
	u.Code = m["code"].(string)
}
