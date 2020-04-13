package account

import (
	"github.com/pkg/errors"
	"log"
)

type AccountService struct {
	AccountRepository AccountRepository
}

func ProvideAccountService(p AccountRepository) AccountService {
	return AccountService{AccountRepository: p}
}

func (s *AccountService) FindAll() []Account {
	return s.AccountRepository.FindAll()
}

func (s *AccountService) FindByID(id string) (Account, error) {
	account, err := s.AccountRepository.FindByID(id)
	if err != nil {
		return Account{}, errors.Wrap(err, "error while finding account")
	}
	return account, err
}

func (s *AccountService) Save(account Account) (Account, error) {
	account, err := s.AccountRepository.Save(account)
	if err != nil {
		return Account{}, errors.Wrap(err, "error while saving account")
	}
	return account, nil
}

func (s *AccountService) Update(account Account) (Account, error) {
	account, err := s.AccountRepository.UpdateName(
		account.ID, []string{account.FirstName, account.LastName})
	if err != nil {
		log.Println(err)
		return Account{}, errors.Wrap(err,"error while updating account")
	}
	return account, nil
}

func (s *AccountService) Delete(account Account) {
	s.AccountRepository.Delete(account)
}
