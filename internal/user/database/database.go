package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//var db *gorm.DB

type User struct {
	ID       int    `json:"id" gorm:"primarily_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func DbConn() *gorm.DB {
	dsn := "root:97216017@tcp(127.0.0.1:3306)/login?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Println("Connection Failed to Open")
	}else{
		log.Println("Connection Established")
		_ = Db.AutoMigrate()
	}
	return Db
}

func (user *User)CreateUser(db *gorm.DB) {
	result := db.Create(user)
	log.Println(result.Error)
}

//var user domain.User
func (u *User)Find(db *gorm.DB) User {
	var user User
	us := &User{Username: u.Username}
	db.Where(us).First(&user)
	return user
}
