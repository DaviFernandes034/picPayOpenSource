package model


type User struct {

	UserId int
	Name string
	LastName string
	Document string
	Email string
	Password string
	Balance float64
	UserType string
}

func (u *User) Subtract(value float64){

	u.Balance -= value


}

func (u *User) Add( value float64){

	u.Balance += value
}
