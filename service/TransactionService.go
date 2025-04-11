package service

import (
	"database/sql"
	"desafio-pic-pay-open-source/repository"
	"fmt"
)

type TransactionService struct {
	TransactionRepository repository.TransactionRepository
}


func (tr *TransactionService) CreateTableTransaction(db *sql.DB)error{

	err:= tr.TransactionRepository.CriaTabelaTransaction(db)
	if err != nil {

		return fmt.Errorf("erro: %v",err.Error())
		
	}

	return nil

}