package repository

import (
	"database/sql"
	"desafio-pic-pay-open-source/model"
	"fmt"
)

type TransactionRepository interface {
	CriaTabelaTransaction(db *sql.DB) error
	Save(model model.Transaction) error
	findTrasnferById(id int) (model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}


func NewtypeRepositoryTransfer(db *sql.DB) TransactionRepository {

	return &transactionRepository{db: db}
}


// Save implements TransactionRepository.
func (t *transactionRepository) Save(model model.Transaction) error {
	panic("unimplemented")
}

// findTrasnferById implements TransactionRepository.
func (t *transactionRepository) findTrasnferById(id int) (model.Transaction, error) {
	panic("unimplemented")
}

// CriaTabelaTransaction implements TransactionRepository.
func (t *transactionRepository) CriaTabelaTransaction(db *sql.DB) error {

	query := `
		
		create table if not exists transfer (
		
			id integer primary key autoincrement,
			amount integer not null,
			senderID integer not null,
			receiverID integer not null,
			localDaetTime text not null,
			foreign key (senderID) references users(id),
			foreign key (receiverID) references users(id)

		);

	`

	_, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("erro ao criar a tabela transfer: %v", err.Error())
	}

	fmt.Println("tabela transfer criada com sucesso")

	return nil
}
