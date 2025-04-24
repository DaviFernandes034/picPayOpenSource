package service

import (
	
	"desafio-pic-pay-open-source/model"
	"desafio-pic-pay-open-source/repository"
	"errors"
	"fmt"
	
)

type UserService struct {

	UserRepository repository.UserRepository


}

func (u *UserService) ValidateTransaction(user model.User, amount float64) error {

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

func (us *UserService) AllUsers(){

	us.UserRepository.UserAll()
}

func (us *UserService) CreateTableUsers()error{

	
	err:= us.UserRepository.CriaTabelaUser()
	if err != nil {

		return fmt.Errorf("erro: %v",err.Error())
		
	}


	return nil
}

func (us *UserService) CreateUser(model model.UserDTO)error{

	err := us.UserRepository.Create(model)
	if err != nil {

		return fmt.Errorf("%s", err.Error())
	}

	return nil
}