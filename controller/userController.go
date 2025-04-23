package controller

import (
	"desafio-pic-pay-open-source/model"
	"desafio-pic-pay-open-source/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}



func (uc *UserController) CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		var request model.UserDTO

		err:= c.BindJSON(&request)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"message":err.Error()})
			return
		}

		
		err = uc.UserService.UserRepository.Create(request)
		if err != nil {
			
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}


		c.Status(http.StatusCreated)


	}
}

func (uc *UserController) GetUserAll() gin.HandlerFunc{

	return func(c *gin.Context) {

		users, err:= uc.UserService.UserRepository.UserAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

