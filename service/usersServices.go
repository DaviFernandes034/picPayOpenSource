package service

import (
	"database/sql"
	"desafio-pic-pay-open-source/model"
	"desafio-pic-pay-open-source/repository"
	"errors"
	"fmt"
	
)

type UserService struct {

	UserRepository repository.UserRepository

}

func validateTransaction(user model.User, amount float64) error {

	if (user.UserType == "lojista"){

		return errors.New("usuario do tipo lojista não pode fazer transação")
	}

	if ( user.Balance < amount ){

		return errors.New("usuario não possui saldo suficiente")
	}

	return nil
}


func (us *UserService) FindUserById(id int){

	us.UserRepository.FindUserById(id)
}

func (us *UserService) Save(model model.User){

	us.UserRepository.Save(model)
}

func (us *UserService) CreateTableUsers(db *sql.DB)error{

	db.Exec("PRAGMA foreign_keys = ON")
	err:= us.UserRepository.CriaTabelaUser(db)
	if err != nil {

		return fmt.Errorf("erro: %v",err.Error())
		
	}


	return nil
}