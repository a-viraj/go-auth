package router

import (
	controller "github.com/a-viraj/golang-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/users/signup", controller.Signup())
	r.POST("/users/login", controller.Login())
}
