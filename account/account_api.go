package account

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AccountAPI struct {
	AccountService AccountService
}

func ProvideAccountAPI(p AccountService) AccountAPI {
	return AccountAPI{AccountService: p}
}

func (p *AccountAPI) FindAll(c *gin.Context) {
	accounts := p.AccountService.FindAll()

	c.JSON(http.StatusOK, gin.H{"accounts": ToAccountModels(accounts)})
}

func (p *AccountAPI) FindByID(c *gin.Context) {
	id :=  c.Param("id")
	account, err := p.AccountService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": ToAccountModel(account)})
}

func (p *AccountAPI) Create(c *gin.Context) {
	var accountModel AccountModel
	err := c.BindJSON(&accountModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := accountModel.Password
	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost) //return []byte
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
	}
	accountModel.Password = string(byteHash)

	account, err := p.AccountService.Save(ToAccount(accountModel))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// Todo find better way to do this
	userPath := location.Get(c).Host+"/api/v1/auth/login"
	c.Writer.Header().Set("Location", userPath)

	ac := ToAccountModel(account)
	ac.Password = ""
	c.JSON(http.StatusOK, gin.H{"account": ac})
}

func (p *AccountAPI) Update(c *gin.Context) {
	var accountNames NamesModel
	err := c.BindJSON(&accountNames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Getting user id from context
	value := c.MustGet("user_id")
	id, ok := value.(string)
	if !ok {
		log.Printf("got data of type %T but wanted int", value)
	}
	account := Account{Base: Base{ID: id}, FirstName: accountNames.FirstName, LastName:accountNames.LastName }
	modifiedAccount, err := p.AccountService.Update(account)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ac := ToAccountModel(modifiedAccount)
	ac.Password = ""
	c.JSON(http.StatusOK, gin.H{"account": ac})
}

func (p *AccountAPI) Delete(c *gin.Context) {
	//id := c.Param("id")
	account := Account{} //p.AccountService.FindByID(id)

	if account == (Account{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.AccountService.Delete(account)

	c.Status(http.StatusOK)
}
