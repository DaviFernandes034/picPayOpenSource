package repository

import (
	"database/sql"
	"fmt"
	
)

//aqui é aonde ficrá as querys do banco de dados em memoria

func CriaTabelaUser(db *sql.DB) error{


	query:= `
			create table if not exists users (
			
				id integer primary key autoincrement,
				name text not null,
				lastName text not null,
				document text not null,
				email text not null,
				password text not null,
				balance integer not null,
				userType text not null,
				unique(document, email)
			);
			`

			_, err:= db.Exec(query)
			if err != nil {

				return fmt.Errorf("erro ao criar a tabela users: %v", err.Error())
			}

			fmt.Println("tabela criada users com sucesso")
			return nil

}


//amanha faço
func CriaTabelaTransaction(db *sql.DB) error{

	query:= `
	
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

	_,err:= db.Query(query)
	if err != nil{
		return fmt.Errorf("erro ao criar a tabela transfer: %v", err.Error())
	}

	fmt.Println("tabela transfer criada com sucesso")

	return nil
}