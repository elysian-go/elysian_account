package account

func ToAccount(accountDTO AccountDTO) Account {
	return Account{}
}

func ToAccountDTO(account Account) AccountDTO {
	return AccountDTO{}
}

func ToAccountDTOs(accounts []Account) []AccountDTO {
	accountdtos := make([]AccountDTO, len(accounts))

	for i, itm := range accounts {
		accountdtos[i] = ToAccountDTO(itm)
	}

	return accountdtos
}