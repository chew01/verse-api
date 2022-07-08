package routes

import (
	"github.com/chew01/verse-api/controllers"
	"github.com/gin-gonic/gin"
)

func (r Routes) User(rg *gin.RouterGroup) {
	userGroup := rg.Group("/user")

	user := new(controllers.UserController)
	userGroup.GET("/", user.GetUsers)
	userGroup.GET("/:name", user.GetUserByName)
	userGroup.POST("/", user.CreateUser)
}
