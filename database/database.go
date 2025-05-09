package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Opening the connection with the database
	stringConnection := "admIndustrial:sysIndustrial@/industrial?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(db)
}
