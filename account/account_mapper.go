package account

func ToAccount(accountModel Model) Account {
	return Account{
		Email: accountModel.Email,
		Password: accountModel.Password,
		FirstName: accountModel.FirstName,
		LastName: accountModel.LastName,
		LastAccess: accountModel.LastAccess,
		Activated: accountModel.Activated,
		AvatarUrl: accountModel.AvatarUrl,
	}
}

func ToAccountModel(account Account) Model {
	return Model{
		ID: account.ID,
		Email: account.Email,
		Password: account.Password,
		FirstName: account.FirstName,
		LastName: account.LastName,
		LastAccess: account.LastAccess,
		Activated: account.Activated,
		AvatarUrl: account.AvatarUrl,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

func ToAccountModels(accounts []Account) []Model {
	accountModels := make([]Model, len(accounts))

	for i, itm := range accounts {
		accountModels[i] = ToAccountModel(itm)
	}

	return accountModels
}