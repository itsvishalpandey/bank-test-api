package main

import (
	config "bank-test-api/Config"
	routes "bank-test-api/Routes"
	"fmt"
)

func main() {
	fmt.Println("Creating API for adding bank.")

	// Connect to database
	config.ConnectDB()

	// Setup router and Routes
	router := routes.SetupRouter()

	// Run Server
	router.Run(":8000")

}
