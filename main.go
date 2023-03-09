package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:mysql@tcp(127.0.0.1)/user?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	fmt.Printf("success =====> 🚀🚀🚀 %v\n")
}
