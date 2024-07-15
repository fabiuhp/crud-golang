package usecases

import (
	"net/http"

	"github.com/fabiuhp/crud-golang/src/configurations/logger"
	"github.com/fabiuhp/crud-golang/src/configurations/validations"
	"github.com/fabiuhp/crud-golang/src/controller/model/req"
	"github.com/fabiuhp/crud-golang/src/model"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userReq.Email,
		userReq.Password,
		userReq.Name,
		userReq.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Usuario criado com sucesso")
	c.String(http.StatusOK, "")
}
