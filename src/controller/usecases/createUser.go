package usecases

import (
	"fmt"

	"github.com/fabiuhp/crud-golang/src/configurations/rest_err"
	"github.com/fabiuhp/crud-golang/src/model/req"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userReq req.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Existem campos incorretos, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userReq)
}
