package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//var db *gorm.DB

type User struct {
	ID       int    `json:"id" gorm:"praimaly_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func DbConn() *gorm.DB {
	dsn := "root:97216017@tcp(127.0.0.1:3306)/login?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Println("Connection Failed to Open")
	}else{
		log.Println("Connection Established")
		db.AutoMigrate()
	}
	return db
}

type mysqlrepo struct {
	Conn *gorm.DB
}
func (m *mysqlrepo)CreateUser(user *User) {
	result := m.Conn.Create(user)
	log.Println(result.Error)
}

//var user domain.User
//func Find(u *domain.User) domain.User {
//	db.Where(u).First(&user)
//	return user
//}
