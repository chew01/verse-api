package controllers

import (
	"github.com/chew01/verse-api/models"
	"github.com/chew01/verse-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

func (a AuthController) Login(c *gin.Context) {
	var form models.UserLoginForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, "Json error")
		return
	}

	hash, err := user.GetHashByEmail(form.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, "User does not exist")
		return
	}

	match := utils.CheckPasswordHash(form.Password, hash)
	if match == true {
		c.JSON(http.StatusOK, "success")
		return
	}

	c.JSON(http.StatusForbidden, "wrong password")
	return
}
