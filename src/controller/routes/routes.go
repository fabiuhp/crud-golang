package routes

import (
	"github.com/fabiuhp/crud-golang/src/controller/usecases"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", usecases.FindUserById)
	r.GET("/getUserByEmail/:userEmail", usecases.FindUserByEmail)
	r.POST("/createUser", usecases.CreateUser)
	r.PUT("/updateUser", usecases.UpdateUser)
	r.DELETE("/deleteUser/:userId", usecases.DeleteUser)
}
