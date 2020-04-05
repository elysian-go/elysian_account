package account

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
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
	id, _ :=  strconv.Atoi(c.Param("id"))
	account := p.AccountService.FindByID(uint(id))

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

	_, err = p.AccountService.Save(ToAccount(accountModel))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func (p *AccountAPI) Update(c *gin.Context) {
	//var accountModel AccountModel
	//err := c.BindJSON(&accountModel)
	//if err != nil {
	//	log.Fatalln(err)
	//	c.Status(http.StatusBadRequest)
	//	return
	//}
	//
	//id, _ :=  strconv.Atoi(c.Param("id"))
	//account := p.AccountService.FindByID(uint(id))
	//if account == (Account{}) {
	//	c.Status(http.StatusBadRequest)
	//	return
	//}
	//
	//account.Code = accountModel.Code
	//account.Price = accountModel.Price
	//p.AccountService.Save(account)

	c.Status(http.StatusOK)
}

func (p *AccountAPI) Delete(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	account := p.AccountService.FindByID(uint(id))
	if account == (Account{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.AccountService.Delete(account)

	c.Status(http.StatusOK)
}
