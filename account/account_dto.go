package account

type AccountDTO struct {
	ID		uint 	`json:"id,string,omitempty"`
	Code	string	`json:"code"`
	Price	uint	`json:"price,string"`
}

