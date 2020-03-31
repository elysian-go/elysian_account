package account

import (
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{"accounts": ToAccountDTOs(accounts)})
}

func (p *AccountAPI) FindByID(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	account := p.AccountService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"account": ToAccountDTO(account)})
}

func (p *AccountAPI) Create(c *gin.Context) {
	var accountDTO AccountDTO
	err := c.BindJSON(&accountDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdAccount := p.AccountService.Save(ToAccount(accountDTO))

	c.JSON(http.StatusOK, gin.H{"account": ToAccountDTO(createdAccount)})
}

func (p *AccountAPI) Update(c *gin.Context) {
	var accountDTO AccountDTO
	err := c.BindJSON(&accountDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ :=  strconv.Atoi(c.Param("id"))
	account := p.AccountService.FindByID(uint(id))
	if account == (Account{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	account.Code = accountDTO.Code
	account.Price = accountDTO.Price
	p.AccountService.Save(account)

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
