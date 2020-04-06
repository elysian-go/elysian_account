package auth

import (
	"errors"
	"github.com/VictorDebray/elysian_account/account"
)

type AuthService struct {
	AccountRepository account.AccountRepository
}

func ProvideAuthService(p account.AccountRepository) AuthService {
	return AuthService{AccountRepository: p}
}

func (p *AuthService) GetByEmail(email string) (account.Account, error) {
	account, err := p.AccountRepository.FindByEmail(email)

	if err != nil {
		return account, errors.New("email does not exist")
	}
	return account, nil
}
