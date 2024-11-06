package main

import (
	"ecommerce-app/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.SetUpRoutes()

	fmt.Println("Starting serwer on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
