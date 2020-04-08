package account

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AccountRepository struct {
	DB *gorm.DB
}

func ProvideAccountRepository(DB *gorm.DB) AccountRepository {
	return AccountRepository{DB: DB}
}

func (r *AccountRepository) FindAll() []Account {
	var accounts []Account
	r.DB.Find(&accounts)

	return accounts
}

func (r *AccountRepository) FindByID(id string) (Account, error) {
	var account Account
	query := *r.DB.Raw("SELECT * FROM account WHERE id = ?", id)
	err := query.Error
	if err != nil {
		return account, errors.New("internal error")
	}

	err = query.Scan(&account).Error
	if err != nil {
		return Account{}, errors.New("internal error")
	}
	return account, nil
}

func (r *AccountRepository) FindByEmail(email string) (Account, error) {
	var account Account
	err := r.DB.First(&account, "email = ?", email).Error
	if err != nil {
		return account, err
	}

	return account, nil
}

func (r *AccountRepository) Save(account Account) (Account, error) {
	err := r.DB.Save(&account).Error
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

func (r *AccountRepository) UpdateName(id string, values []string) (Account, error) {
	err := r.DB.Exec("UPDATE account SET (first_name, last_name) = (?) WHERE id = ?", values, id).Error
	if err != nil {
		return Account{}, err
	}
	return r.FindByID(id)
}

func (r *AccountRepository) Delete(account Account) {
	r.DB.Delete(&account)
}
