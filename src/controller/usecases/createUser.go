package usecases

import (
	"fmt"
	"log"

	"github.com/fabiuhp/crud-golang/src/configurations/validations"
	"github.com/fabiuhp/crud-golang/src/model/req"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Iniciando controller: CreateUser")
	var userReq req.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		log.Printf("Erro tentando marshal o objeto, error=%s\n", err.Error())
		errRest := validations.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userReq)
}
