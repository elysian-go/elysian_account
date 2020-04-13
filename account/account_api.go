package account

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

type AccountAPI struct {
	AccountService AccountService
}

func ProvideAccountAPI(p AccountService) AccountAPI {
	return AccountAPI{AccountService: p}
}

func (a *AccountAPI) FindAll(c *gin.Context) {
	accounts := a.AccountService.FindAll()

	c.JSON(http.StatusOK, gin.H{"accounts": ToAccountModels(accounts)})
}

func (a *AccountAPI) FindByID(c *gin.Context) {
	id := c.Param("id")
	account, err := a.AccountService.FindByID(id)
	if err != nil {
		log.Println(err)
		switch {
		case gorm.IsRecordNotFoundError(errors.Cause(err)):
			c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": ToAccountModel(account)})
}

func (a *AccountAPI) Create(c *gin.Context) {
	var accountModel Model
	err := c.BindJSON(&accountModel)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := accountModel.Password
	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost) //return []byte
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
	}
	accountModel.Password = string(byteHash)

	account, err := a.AccountService.Save(ToAccount(accountModel))
	if err != nil {
		log.Println(err)
		switch {
		case strings.Contains(err.Error(), "duplicate"):
			c.JSON(http.StatusConflict, gin.H{"error": "duplicate email"})
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	// Todo find better way to do this
	userPath := location.Get(c).Host + "/api/v1/auth/login"
	c.Writer.Header().Set("Location", userPath)

	ac := ToAccountModel(account)
	ac.Password = ""
	c.JSON(http.StatusOK, gin.H{"account": ac})
}

func (a *AccountAPI) Update(c *gin.Context) {
	var accountNames NamesModel
	err := c.BindJSON(&accountNames)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Getting user id from context
	value := c.MustGet("user_id")
	id, ok := value.(string)
	if !ok {
		log.Printf("got data of type %T but wanted string", value)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
	account := Account{Base: Base{ID: id}, FirstName: accountNames.FirstName, LastName: accountNames.LastName}
	modifiedAccount, err := a.AccountService.Update(account)
	if err != nil {
		log.Println(err)
		switch {
		case gorm.IsRecordNotFoundError(errors.Cause(err)):
			c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	ac := ToAccountModel(modifiedAccount)
	ac.Password = ""
	c.JSON(http.StatusOK, gin.H{"account": ac})
}

func (a *AccountAPI) Delete(c *gin.Context) {
	//id := c.Param("id")
	account := Account{} //a.AccountService.FindByID(id)

	if account == (Account{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	a.AccountService.Delete(account)

	c.Status(http.StatusOK)
}
