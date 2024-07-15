package service

import (
	"fmt"

	"github.com/fabiuhp/crud-golang/src/configurations/logger"
	"github.com/fabiuhp/crud-golang/src/configurations/rest_err"
	"github.com/fabiuhp/crud-golang/src/model"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Iniciando criação de modelo")

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
