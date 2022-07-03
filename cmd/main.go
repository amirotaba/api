package main

import (
	"login/internal/user/database"
	"login/internal/user/deliver/httpd"
)

func main() {
	database.DbConn()
	httpd.Deliver()
}