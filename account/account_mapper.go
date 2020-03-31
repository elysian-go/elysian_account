package account

func ToAccount(accountDTO AccountDTO) Account {
	return Account{Code: accountDTO.Code, Price: accountDTO.Price}
}

func ToAccountDTO(account Account) AccountDTO {
	return AccountDTO{ID: account.ID, Code: account.Code, Price: account.Price}
}

func ToAccountDTOs(accounts []Account) []AccountDTO {
	accountdtos := make([]AccountDTO, len(accounts))

	for i, itm := range accounts {
		accountdtos[i] = ToAccountDTO(itm)
	}

	return accountdtos
}