package main

import (
	"github.com/user/go-curd/initializers"
	"github.com/user/go-curd/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
