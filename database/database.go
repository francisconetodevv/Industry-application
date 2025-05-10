package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/*
This function it is responsible to stablish the connection with the database
*/
func Connection() (*sql.DB, error) {
	// Opening the connection with the database
	stringConnection := "admIndustrial:sysIndustrial@/industrial?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		return nil, erro
	}

	/*
		db.Ping() - It evaluates the connection with the database: May be used to determine if communication with the database server is still possible
	*/
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
