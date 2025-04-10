package main

import (
	"desafio-pic-pay-open-source/repository"
	"github.com/gin-gonic/gin"
	
	"log"
)

func main(){

	router:= gin.Default()

	db, err:= repository.Init()
	if err != nil {

		log.Fatal("erro ao abrir o banco de dados")
	}

	defer db.Close()
	db.Exec("PRAGMA foreign_keys = ON")

	err = repository.CriaTabelaUser(db)
	if err != nil {

		log.Fatal(err)
	}

	router.Run(":8080")

}