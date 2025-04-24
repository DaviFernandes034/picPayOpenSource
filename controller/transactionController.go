package controller

import (
	"desafio-pic-pay-open-source/model"
	"desafio-pic-pay-open-source/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {

	TranferController service.TransactionService
}


func (tc *TransactionController) CreateTransfer() gin.HandlerFunc {


	return func(c *gin.Context) {

		var request model.TransactionDto

		err:= c.BindJSON(&request)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err= tc.TranferController.CreateTransaction(request)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"messsage": err.Error()})
		}


		c.Status(http.StatusCreated)

	}
}


func (tc *TransactionController) GetAllTransfer() gin.HandlerFunc {


	return func(c *gin.Context) {

		tranfers, err:= tc.TranferController.AllTranferGet()
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}


		c.JSON(http.StatusOK, tranfers)

	}
}