package auth

type AuthModel struct {
	Email			string			`json:"email" binding:"required,email"`
	Password		string			`json:"password" binding:"required,max=16,min=8"`
}

