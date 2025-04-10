package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//iniciando o banco de dados em memoria
func Init()(*sql.DB, error){


	db, err:= sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}