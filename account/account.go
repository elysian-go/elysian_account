package account

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Code string
	Price uint
}
