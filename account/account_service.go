package account

import (
	"errors"
)

type AccountService struct {
	AccountRepository AccountRepository
}

func ProvideAccountService(p AccountRepository) AccountService {
	return AccountService{AccountRepository: p}
}

func (p *AccountService) FindAll() []Account {
	return p.AccountRepository.FindAll()
}

func (p *AccountService) FindByID(id uint) Account {
	return p.AccountRepository.FindByID(id)
}

func (p *AccountService) Save(account Account) (Account, error) {
	account, err := p.AccountRepository.Save(account)
	if err != nil {
		return account, errors.New("duplicate entry on email")
	}
	return account, nil
}

func (p *AccountService) Delete(account Account) {
	p.AccountRepository.Delete(account)
}

