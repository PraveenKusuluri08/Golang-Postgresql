package main

import (
	"github.com/Praveenkusuluri08/helpers"
	"github.com/Praveenkusuluri08/models"
)

func init() {
	helpers.LoadEnvVariables()
	helpers.DBConnection()
}
func main() {
	helpers.DB.AutoMigrate(models.Posts{})
}
