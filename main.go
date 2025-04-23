package main

import (
	"desafio-pic-pay-open-source/controller"
	"desafio-pic-pay-open-source/repository"
	"desafio-pic-pay-open-source/service"

	"github.com/gin-gonic/gin"

	"log"
)

func main(){

	Router:= gin.Default()

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
	//controller User

	uc:= controller.UserController{
		UserService: us,
	}


	//criação da tabela de usuarios
	err = us.CreateTableUsers()
	if err != nil {

		log.Fatal(err)
	}
	//criaçao da tabela de trasnferencia
	err = ts.CreateTableTransaction()
	if err != nil{
		log.Fatal(err)
	}


	Router.POST("/create", uc.CreateUser())
	Router.GET("/AllUsers", uc.GetUserAll())
	
		


	Router.Run(":8080")

}

