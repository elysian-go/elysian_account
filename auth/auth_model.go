package auth

import "github.com/VictorDebray/elysian_account/account"

type AuthModel struct {
	Email			string			`json:"email" binding:"required,email"`
	Password		string			`json:"password" binding:"required,max=16,min=8"`
}

func ToAuthModel(account account.Account) AuthModel {
	return AuthModel{
		Email: account.Email,
		Password: account.Password,
	}
}