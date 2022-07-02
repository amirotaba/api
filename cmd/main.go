package main

import (
	"login/internal/user/database"
	"login/internal/user/delivery/httpd"
)

func main() {
	httpd.Deliver()
	database.DbConn()
}