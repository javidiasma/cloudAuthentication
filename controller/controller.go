package controller

import (
	"cloudAuthentication/DAL"
	"cloudAuthentication/DAO"
	"cloudAuthentication/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InputToken struct {
	Token string `json:"token"`
}

func SignUp(c *gin.Context) {
	var input DAO.UserRequestModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		fmt.Println(111111)
		return
	}
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "enter password"})
		return
	}
	if input.Username == "" {
		fmt.Println("asma")
		c.JSON(http.StatusBadRequest, gin.H{"status": "enter username"})
		return
	}
	err := DAL.UsernameExists(input.Username)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"status": "user already exists"})
		return
	}
	token, err := utils.CreateTokens(input.Username, input.Password)
	fmt.Println(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	_, err2 := DAL.SignUp(input)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": token})
	return
}

func ValidateUser(c *gin.Context) {
	var input InputToken
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		fmt.Println(111111)
		return
	}
	exists, _ := utils.ValidateToken(input.Token)
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"status": "unauthorized3"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "validated"})
	return
}
