package main

import (
	"io"
	"os"

	"github.com/Praveenkusuluri08/helpers"
	"github.com/Praveenkusuluri08/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()
	helpers.DBConnection()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.LoggerWithFormatter(helpers.LogParse))
	r.Use(gin.Recovery())

	routes.PostRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
