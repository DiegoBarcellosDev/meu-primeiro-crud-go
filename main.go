package main

import (
	"log"

	"github.com/DiegoBarcellosDev/meu-primeiro-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}