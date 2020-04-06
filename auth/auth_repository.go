package auth

import (
	//"github.com/VictorDebray/elysian_account/account"
	"github.com/jinzhu/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func ProvideAuthRepostiory(DB *gorm.DB) AuthRepository {
	return AuthRepository{DB: DB}
}
