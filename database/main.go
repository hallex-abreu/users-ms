package database

import (
	"github.com/hallex-abreu/users-ms/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "user"
const DB_PASSWORD = "user"
const DB_NAME = "userclientes"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB

func Connection() {
	dns := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&entities.Users{})
	db.AutoMigrate(&entities.Roles{})
	db.AutoMigrate(&entities.UsersRoles{})
	db.AutoMigrate(&entities.Permissions{})
	db.AutoMigrate(&entities.RolesPermissions{})

	DB = db
}
