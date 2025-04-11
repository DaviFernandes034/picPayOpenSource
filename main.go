package main

import (
	"desafio-pic-pay-open-source/repository"
	"github.com/gin-gonic/gin"
	"desafio-pic-pay-open-source/service"
	
	"log"
)

func main(){

	router:= gin.Default()

	db, err:= repository.Init()
	if err != nil {

		log.Fatal("erro ao abrir o banco de dados")
	}
	defer db.Close()
	
	//repositorios de users (ru) e transaction (rt)
	ru:= repository.NewtypeRepository(db)
	rt:= repository.NewtypeRepositoryTransfer(db)

	us:= service.UserService{ //services de users
		 UserRepository: ru,
	}
	ts:= service.TransactionService{ //services de transaction
		TransactionRepository: rt,
	}



	//criação da tabela de usuarios
	err = us.CreateTableUsers(db)
	if err != nil {

		log.Fatal(err)
	}
	//criaçao da tabela de trasnferencia
	err = ts.CreateTableTransaction(db)
	if err != nil{
		log.Fatal(err)
	}



	router.Run(":8080")

}