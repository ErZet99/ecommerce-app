package routes

import (
	"ecommerce-app/controllers"
	"net/http"
)

func SetUpRoutes() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/products", controllers.GetProducts)
	http.HandleFunc("auth/register", controllers.Register)
	http.HandleFunc("/auth/login", controllers.Login)
}
