package auth

import "github.com/VictorDebray/elysian_account/account"

type Model struct {
	Email			string			`json:"email" binding:"required,email"`
	Password		string			`json:"password" binding:"required,max=16,min=8"`
}

func ToAuthModel(account account.Account) Model {
	return Model{
		Email: account.Email,
		Password: account.Password,
	}
}