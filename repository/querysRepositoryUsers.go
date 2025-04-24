package repository

import (
	"database/sql"
	"desafio-pic-pay-open-source/model"
	"fmt"

)

type UserRepository interface {
	CriaTabelaUser() error
	FindUserById(id int) (model.User, error)
	Create(model model.UserDTO) error
	Save(model model.User) error
	UserAll() ([]model.User, error)
}

type userRepository struct {
	db *sql.DB
}


func NewtypeRepository(db *sql.DB) UserRepository {

	return &userRepository{db: db}
}

// UserAll implements UserRepository.
func (u *userRepository) UserAll() ([]model.User, error) {

	query:= "select id, name, lastName,document, email, password, balance, userType from users"

	rows, err:= u.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro na query: %w", err)
	}

	defer rows.Close()

	var users []model.User

	for rows.Next(){

		var user model.User

		if err:= rows.Scan(
				
			&user.UserId,
			&user.Name,
			&user.LastName,
			&user.Document,
			&user.Email,
			&user.Password,
			&user.Balance,
			&user.UserType); err != nil {

				return nil, fmt.Errorf("erro ao passar os dados para a struct user, %w", err)
			}

			users = append(users, user)
	}

	if err:= rows.Err(); err != nil {
		return nil, fmt.Errorf("erro entre as iterações: %w", err)
	}


	return users, nil
}



// Save implements UserRepository.
func (u *userRepository) Save(model model.User) error {

	query := `update users
			set balance = ?
			where id = ?`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("erro ao preparar a query; %v", err)
	}

	_, err = stmt.Exec(model.Balance, model.UserId)
	if err != nil {
		return fmt.Errorf("erro ao executar o stmt: %v", err)
	}

	return nil

}

// Create implements UserRepository.
func (u *userRepository) Create(model model.UserDTO) error {

	query := `insert into users(name, lastName,document, email, password, balance, userType)
			  values (?,?,?,?,?,?,?)`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("erro ao preparar a query; %v", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(model.Name, model.LastName, model.Document, model.Email, model.Password, model.Balance, model.UserType)
	if err != nil {
		return fmt.Errorf("erro ao executar o stmt: %v", err)
	}

	return nil
}

// FindUserById implements UserRepository.
func (u *userRepository) FindUserById(id int) (model.User, error) {

	query := `select id, name, lastName,document, email, password, balance, userType from users 
				where id = ?`

	var user model.User

	err := u.db.QueryRow(query, id).Scan(

		&user.UserId,
		&user.Name,
		&user.LastName,
		&user.Document,
		&user.Email,
		&user.Password,
		&user.Balance,
		&user.UserType,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("usuario com id %d não encontrado", id)
		}

		return model.User{}, fmt.Errorf("erro ao buscar usuario: %v", err)
	}

	return user, nil

}

//aqui é aonde ficrá as querys do banco de dados em memoria

func (u *userRepository) CriaTabelaUser() error {

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

	_, err := u.db.Exec(query)
	if err != nil {

		return fmt.Errorf("erro ao criar a tabela users: %v", err.Error())
	}

	fmt.Println("tabela criada users com sucesso")
	return nil

}
