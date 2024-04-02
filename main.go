package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Env.DBUser,
		Passwd:               Env.DBPassword,
		Addr:                 Env.DBAddress,
		DBName:               Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStorage(db)
	api := NewAPIServer(":8080", store)

	api.Start()
}
