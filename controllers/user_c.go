package controllers

import (
	"github.com/chew01/verse-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type UserController struct{}

var user = new(models.User)

func (u UserController) GetUsers(c *gin.Context) {
	users, err := user.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Database error")
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u UserController) GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := user.GetOneByName(name)
	if err == pgx.ErrNoRows {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Database error")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UserController) CreateUser(c *gin.Context) {
	var form models.UserSignupForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, "Json error")
		return
	}

	if err := user.Create(form); err != nil {
		c.JSON(http.StatusInternalServerError, "Database error")
		return
	}

	c.JSON(http.StatusCreated, "Created")
}
