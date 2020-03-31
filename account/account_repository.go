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

func (p *AccountRepository) Save(account Account) Account {
	p.DB.Save(&account)

	return account
}

func (p *AccountRepository) Delete(account Account) {
	p.DB.Delete(&account)
}

