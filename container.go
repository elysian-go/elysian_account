//+build wireinject

package main

import (
	"github.com/VictorDebray/elysian_account/account"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitAccountAPI(db *gorm.DB) account.AccountAPI {
	panic(wire.Build(
		account.ProvideAccountRepostiory, account.ProvideAccountService, account.ProvideAccountAPI))
}