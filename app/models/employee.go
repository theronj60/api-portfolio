package models

import (
	"github.com/theronj60/api-portfolio/app"
	"gorm.io/gorm"
)

var db *gorm.DB

type Employee struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	app.Connect()
	db = app.GetDB()
	db.AutoMigrate(&Employee{})
}
