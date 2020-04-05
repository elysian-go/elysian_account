package account

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AccountRepository struct {
	DB *gorm.DB
}

func ProvideAccountRepostiory(DB *gorm.DB) AccountRepository {
	return AccountRepository{DB: DB}
}

func (p *AccountRepository) FindAll() []Account {
	var accounts []Account
	p.DB.Find(&accounts)

	return accounts
}

func (p *AccountRepository) FindByID(id uint) Account {
	var account Account
	p.DB.First(&account, id)

	return account
}

func (p *AccountRepository) FindByEmail(email string) (Account, error) {
	var account Account
	err := p.DB.First(&account, "email = ?", email).Error
	if err != nil {
		return account, err
	}

	return account, nil
}

func (p *AccountRepository) Save(account Account) (Account, error) {
	err := p.DB.Save(&account).Error
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

func (p *AccountRepository) Delete(account Account) {
	p.DB.Delete(&account)
}
