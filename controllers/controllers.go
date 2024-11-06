package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the E-Commerce API")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "List of products")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Register a new user")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Login user")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
