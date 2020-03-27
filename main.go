package main

import (
	"JWT-Based-Authentication/server"
	"JWT-Based-Authentication/utils"
	"fmt"
)

func main() {

	_, err := utils.NewDatabase()

	if err != nil {
		fmt.Print(err)
	}

	//models.CreateTable()

	server.Run(":5000")
}
