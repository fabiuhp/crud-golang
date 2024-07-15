package usecases

import (
	"fmt"
	"net/http"

	"github.com/fabiuhp/crud-golang/src/configurations/logger"
	"github.com/fabiuhp/crud-golang/src/configurations/validations"
	"github.com/fabiuhp/crud-golang/src/model/req"
	"github.com/fabiuhp/crud-golang/src/model/res"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	logger.Info("Iniciando controller: CreateUser")
	var userReq req.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		logger.Error("Erro tentando marshal o objeto", err)
		errRest := validations.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userReq)
	res := res.UserResponse{
		ID:    "teste",
		Email: userReq.Email,
		Name:  userReq.Name,
		Age:   userReq.Age,
	}
	logger.Info("Usuario criado com sucesso")
	c.JSON(http.StatusOK, res)
}
