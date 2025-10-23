package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // mux is a router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products",
		http.HandlerFunc(handlers.CreateProducts))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))
	fmt.Println("Server is running on port 3000")
	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}