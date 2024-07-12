package usecases

import (
	"github.com/fabiuhp/crud-golang/src/configurations/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamou a rota errada")
	c.JSON(err.Code, err)
}
