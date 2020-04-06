package account

import (
	"time"
)

type AccountModel struct {
	ID				string			`json:"id,string"`
	Email			string			`json:"email" binding:"required,email"`
	// ",omitempty" enables to strip password when returning struct
	Password		string			`json:",omitempty" binding:"required,max=16,min=8"`
	FirstName		string			`json:"first_name, omitempty" binding:"required,alpha,max=15"`
	LastName		string			`json:"last_name, omitempty" binding:"required,alpha,max=15"`
	LastAccess		time.Time		`json:"-"`
	Activated		bool			`json:"-"`
	AvatarUrl		string			`json:"avatar_url,omitempty"`
	CreatedAt		time.Time		`json:"created_at,omitempty"`
	UpdatedAt		time.Time		`json:"updated_at,omitempty"`
}

type NamesModel struct {
	FirstName		string			`json:"first_name, omitempty" binding:"required,alpha,max=15"`
	LastName		string			`json:"last_name, omitempty" binding:"required,alpha,max=15"`
}