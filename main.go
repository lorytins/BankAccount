package main

import (
	"fmt"

	"github.com/lorytins/BankAcountApi/database"
	"github.com/lorytins/BankAcountApi/routes"
)

func main() {
	fmt.Println("Starting server on the port 8080...")
	database.ConnectWithDataBase()
	routes.HandleRequest()
}