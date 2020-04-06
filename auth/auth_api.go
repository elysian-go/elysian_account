package auth

import (
	"github.com/gin-contrib/location"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthAPI struct {
	AuthService AuthService
}

func ProvideAuthAPI(p AuthService) AuthAPI {
	return AuthAPI{AuthService: p}
}

func (p *AuthAPI) Login(c *gin.Context) {
	var authModel Model
	err := c.BindJSON(&authModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	account, err := p.AuthService.GetByEmail(authModel.Email)
	bytePassword :=	[]byte(authModel.Password)
	byteHash := []byte(account.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "invalid login"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", account.ID)
	session.Save()

	// Todo find better way to do this
	userPath := location.Get(c).Host+"/api/v1/account"
	c.Writer.Header().Set("Location", userPath)

	c.Status(http.StatusNoContent)
}

func (p *AuthAPI) Logout(c* gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Account logout"})
}