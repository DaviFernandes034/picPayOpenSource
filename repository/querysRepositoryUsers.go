package repository

import (
	"database/sql"
	"desafio-pic-pay-open-source/model"
	"fmt"
)

type UserRepository interface {
	CriaTabelaUser(db *sql.DB) error
	FindUserById(id int) (model.User, error)
	Save(model model.User) error
}

type userRepository struct {
	db *sql.DB
}



func NewtypeRepository(db *sql.DB) UserRepository {

	return &userRepository{db: db}
}
// Save implements UserRepository.
func (u *userRepository) Save(model model.User) error {
	panic("unimplemented")
}

// FindUserById implements UserRepository.
func (u *userRepository) FindUserById(id int) (model.User, error) {
	panic("unimplemented")
}

//aqui é aonde ficrá as querys do banco de dados em memoria

func (u *userRepository) CriaTabelaUser(db *sql.DB) error {

	query := `
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

	_, err := db.Exec(query)
	if err != nil {

		return fmt.Errorf("erro ao criar a tabela users: %v", err.Error())
	}

	fmt.Println("tabela criada users com sucesso")
	return nil

}


