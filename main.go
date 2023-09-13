package main

import (
	"log"
	"os"

	routes "github.com/a-viraj/golang-auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRouter(router)
	routes.UserRouter(router)
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"Sucess": "access granted to api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"Sucess": "access granted to api-2"})
	})
	router.Run(":" + port)
}
