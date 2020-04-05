package account

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	ID			string		`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DeletedAt	*time.Time	`sql:"index"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type Account struct {
	Base
	Email				string		`gorm:"type:varchar(100);unique_index"`
	Password			string
	FirstName			string
	LastName			string
	LastAccess			time.Time
	Activated			bool
	AvatarUrl			string
	UserType			string
}

func (m *Base) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (m *Base) BeforeCreate(scope *gorm.Scope) error {
	if m.UpdatedAt.IsZero() {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}

	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}