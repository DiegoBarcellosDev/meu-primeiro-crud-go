package controller

import (
	"fmt"

	"github.com/DiegoBarcellosDev/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/DiegoBarcellosDev/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintln("there are some incorrect fields, error=%s", err.Error))
		
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}