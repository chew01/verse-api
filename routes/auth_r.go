package routes

import (
	"github.com/chew01/verse-api/controllers"
	"github.com/gin-gonic/gin"
)

func (r Routes) Auth(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")

	auth := new(controllers.AuthController)
	authGroup.POST("/", auth.Login)
}
