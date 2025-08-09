package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"m-banking.info/model"
)

func CreateAccountHandler(c *gin.Context) {
	var account model.M_Account
	if err := c.ShouldBindingJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
