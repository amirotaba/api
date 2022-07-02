package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"login/internal/domain"
)
var db *gorm.DB

func DbConn() {
	sqlDB, err := sql.Open("mysql", "login")
	DB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err!=nil{
		log.Println("Connection Failed to Open")
	}else{
		log.Println("Connection Established")
		DB.AutoMigrate()
	}
}


var user domain.User
func Find(u *domain.User) domain.User {
	db.Where(u).First(&user)
	return user
}