package routes

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func New() Routes {
	r := Routes{router: gin.Default()}

	v1 := r.router.Group("/v1")
	r.User(v1)
	r.Auth(v1)

	return r
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
