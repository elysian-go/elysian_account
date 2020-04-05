package account

import (
	"time"
)

type AccountModel struct {
	ID				string			`json:"id,string"`
	Email			string			`json:"email" binding:"required,email"`
	Password		string			`json:"password" binding:"required,max=16,min=8"`
	FirstName		string			`json:"first_name, omitempty" binding:"required,alpha,max=15"`
	LastName		string			`json:"last_name, omitempty" binding:"required,alpha,max=15"`
	LastAccess		time.Time		`json:"-"`
	Activated		bool			`json:"-"`
	AvatarUrl		string			`json:"avatar_url, omitempty"`
	CreatedAt		time.Time		`json:"-"`
	UpdatedAt		time.Time		`json:"-"`
}
