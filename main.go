package main

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/router"
)

func main() {
	fmt.Println("Starting...")
	database.DBInit()
	res := helpers.HashPassword("bagus1")
	fmt.Println("password is", res)

	respass := helpers.ComparePassword("bagus1", res)
	fmt.Println("hasil dari hashed password:", respass)
	r := router.StartApp()

	r.Run("127.0.0.1:3000")
}
