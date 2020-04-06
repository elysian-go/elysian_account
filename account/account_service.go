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

func (s *AccountService) FindAll() []Account {
	return s.AccountRepository.FindAll()
}

func (s *AccountService) FindByID(id string) (Account, error) {
	account, err := s.AccountRepository.FindByID(id)
	if err != nil {
		return account, errors.New("resource not found")
	}
	return account, err
}

func (s *AccountService) Save(account Account) (Account, error) {
	account, err := s.AccountRepository.Save(account)
	if err != nil {
		return account, errors.New("duplicate entry on email")
	}
	return account, nil
}

func (s *AccountService) Update(account Account) (Account, error) {
	account, err := s.AccountRepository.UpdateName(
		account.ID, []string{account.FirstName, account.LastName})
	if err != nil {
		return account, errors.New("account update failed")
	}
	return account, nil
}

func (s *AccountService) Delete(account Account) {
	s.AccountRepository.Delete(account)
}
