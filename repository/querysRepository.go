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
				userType text not null
			);
			`

			_, err:= db.Exec(query)
			if err != nil {

				return fmt.Errorf("erro ao criar a tabela: %v", err.Error())
			}

			fmt.Println("tabela criada com sucesso")
			return nil

}


//amanha faço
func CriaTabelaTransaction(db *sql.DB) error{


	return nil
}