package repository

import (
	"database/sql"
	"desafio-pic-pay-open-source/model"
	"fmt"
)

type TransactionRepository interface {
	CriaTabelaTransaction() error
	Create(model model.Transaction) error
	findTransferById(id int) (model.Transaction, error)
	Save(model model.Transaction) error
	AllTranfer() ([]model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}



// Save implements TransactionRepository.
func (t *transactionRepository) Save(model model.Transaction) error {

	panic("yuyfd")
}

func NewtypeRepositoryTransfer(db *sql.DB) TransactionRepository {

	return &transactionRepository{db: db}
}

// AllTranfer implements TransactionRepository.
func (t *transactionRepository) AllTranfer() ([]model.Transaction, error) {
	
	query:= `select id, amount, senderID, receiverID, localDateTime from transfer `
	
	rows, err:= t.db.Query(query)
	if err != nil {

		return nil, fmt.Errorf("erro na query: %w", err)
	}

	defer rows.Close()

	var tranfers []model.Transaction

	for rows.Next() {

		var tranfer model.Transaction

		if err:= rows.Scan(

			&tranfer.TransferID,
			&tranfer.Amount,
			&tranfer.SenderID,
			&tranfer.ReceiverID,
			&tranfer.LocaldateTime); err != nil {

				return nil, fmt.Errorf("erro ao passar os dados para a struct transactio: %w", err)

			}

			tranfers = append(tranfers, tranfer)
	}
	
	if err:= rows.Err(); err != nil {
		return nil, fmt.Errorf("erro entre as iterações: %w", err)
	}

	return tranfers , nil

}


// Save implements TransactionRepository.
func (t *transactionRepository) Create(model model.Transaction) error {

	query := `insert into transfer(amount, senderID, receiverID, localDateTime) values (? , ? , ? ,?)`

	stmt, err := t.db.Prepare(query)
	if err != nil {

		return fmt.Errorf("erro ao preparar o stmt: %v", err)
	}

	_, err = stmt.Exec(model.Amount, model.SenderID, model.ReceiverID, model.LocaldateTime)
	if err != nil {

		return fmt.Errorf("erro ao executar a query: %v", err)
	}

	return nil
}

// findTrasnferById implements TransactionRepository.
func (t *transactionRepository) findTransferById(id int) (model.Transaction, error) {
	panic("unimplemented")
}

// CriaTabelaTransaction implements TransactionRepository.
func (t *transactionRepository) CriaTabelaTransaction() error {

	query := `
		
		create table if not exists transfer (
		
			id integer primary key autoincrement,
			amount integer not null,
			senderID integer not null,
			receiverID integer not null,
			localDateTime text not null,
			foreign key (senderID) references users(id),
			foreign key (receiverID) references users(id)

		);

	`

	_, err := t.db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar a tabela transfer: %v", err.Error())
	}

	fmt.Println("tabela transfer criada com sucesso")

	return nil
}
