package router

import (
	controller "github.com/a-viraj/golang-auth/controllers"
	middleware "github.com/a-viraj/golang-auth/middleware"
	"github.com/gin-gonic/gin"
)
func UserRouter(r *gin.Engine){
	r.Use(middleware.Authenticate())
	r.GET("/users",controller.GetUsers())
	r.GET("/users/:userId",controller.GetUser())
}