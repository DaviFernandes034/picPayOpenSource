package service

import (
	
	"desafio-pic-pay-open-source/model"
	"desafio-pic-pay-open-source/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Structs que representam a resposta
type ResponseData struct { 
	Status string `json:"status"`
	Data   struct {
		Authorization bool `json:"authorization"`
	} `json:"data"`
}

type TransactionService struct {
	TransactionRepository repository.TransactionRepository
	Service UserService
}


func (tr *TransactionService) CreateTableTransaction()error{

	err:= tr.TransactionRepository.CriaTabelaTransaction()
	if err != nil {

		return fmt.Errorf("erro: %v",err.Error())
		
	}

	return nil

}

	
func (tr *TransactionService) CreateTransaction(transaction model.TransactionDto )error{

	sender, err:= tr.Service.UserRepository.FindUserById(transaction.SenderID)
	if err != nil {

		return err
	}

	receiver, err:= tr.Service.UserRepository.FindUserById(transaction.ReceiverID)
	if err != nil {
		return err
	}

	tr.Service.ValidateTransaction(sender, transaction.Amount)


	IsAuthorize, err:= AuthorizeTransaction(sender, transaction.Amount)
	if err != nil {

		return err
	}

	if !IsAuthorize {

		return fmt.Errorf("transação não foi autorizada")
	}


	TransactionDto:= model.TransactionDto{
		Amount: transaction.Amount,
		SenderID: sender.UserId,
		ReceiverID: receiver.UserId,
		

	}

	transactionRequest:= model.Transaction{
		Amount: TransactionDto.Amount,
		SenderID: TransactionDto.SenderID,
		ReceiverID: TransactionDto.ReceiverID,
		LocaldateTime: time.Now(),


	}


	sender.Subtract(transaction.Amount)
	receiver.Add(transaction.Amount)
	err = tr.TransactionRepository.Create(transactionRequest)
	if err != nil {

		return err
	}

	//atualizando o balance dos users
	err= tr.Service.UserRepository.Save(sender)
	if err != nil {

		return err
	}

	err = tr.Service.UserRepository.Save(receiver)
	if err != nil {
		return err
	}




	return nil
}


func AuthorizeTransaction( sender model.User, value float64)(bool, error){

	
	url:="https://util.devi.tools/api/v2/authorize"  //api que retorna se a autenticação deu certo ou nao

	resp, err:= http.Get(url)
	if err != nil {

		return false,fmt.Errorf("erro na resquisição: %w", err)
		
	}
	defer resp.Body.Close()

	body, err:= io.ReadAll(resp.Body)
	if err != nil {
	   return false,fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
		
	}

	var result ResponseData

	if err := json.Unmarshal(body, &result); err != nil {
		return false, fmt.Errorf("erro ao decodificar JSON: %w", err)
		
	}

	if result.Status == "success"{

		return true, nil
	}else {
		return false, nil
	}


}


func (tr *TransactionService) AllTranferGet()([]model.TransactionResponse, error){

	transfer, err:= tr.TransactionRepository.AllTranfer()
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}


	var response []model.TransactionResponse

	for _, t := range transfer {

		sender, err:= tr.Service.UserRepository.FindUserById(t.SenderID)
		if err != nil {

			return nil, fmt.Errorf("%s", err.Error())
		}

		receiver, err:= tr.Service.UserRepository.FindUserById(t.ReceiverID)
		if err != nil {

			return nil, fmt.Errorf("%s", err.Error())
		}

		transaction:= model.TransactionResponse{

			TransferID: t.TransferID,
			Amount: t.Amount,
			Sender: sender,
			Receiver: receiver,
			LocaldateTime: t.LocaldateTime,
		}

		response = append(response, transaction)
	}

	
	return response, nil
	
}

