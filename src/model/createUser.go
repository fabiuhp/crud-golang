package model

import (
	"github.com/fabiuhp/crud-golang/src/configurations/logger"
	"github.com/fabiuhp/crud-golang/src/configurations/rest_err"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Iniciando criação de modelo")

	ud.EncryptPassword()

	return nil
}
