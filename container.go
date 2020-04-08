//+build wireinject

package main

import (
	"github.com/VictorDebray/elysian_account/account"
	"github.com/VictorDebray/elysian_account/auth"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitAccountAPI(db *gorm.DB) account.AccountAPI {
	panic(wire.Build(
		account.ProvideAccountRepository, account.ProvideAccountService, account.ProvideAccountAPI))
}

func InitAuthAPI(db *gorm.DB) auth.AuthAPI {
	panic(wire.Build(
		account.ProvideAccountRepository, auth.ProvideAuthService, auth.ProvideAuthAPI))
}