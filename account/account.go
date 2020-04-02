package account

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID         string		`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt  time.Time	`json:"created_at"`
	UpdatedAt  time.Time	`json:"update_at"`
	DeletedAt *time.Time	`sql:"index" json:"deleted_at"`
}

type Account struct {
	Base
	FirstName			string
	LastName			string
	Email				string
	Password			string
	LastAccess			time.Time
	Activated			bool
	AvatarUrl			string
	UserType			string
}

//todo make available to Base
func (user *Account) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}