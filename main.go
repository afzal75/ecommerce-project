package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

 func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello World!")
 }
 func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I am Golang")
 }

 type Product struct {
	 ID int `json:"id"`
	 Title string `json:"title"`
	 Description string `json:"description"`
	 Price float64 `json:"price"`
	 ImgURL string `json:"imageUrl"`
 }

 var productList []Product

 func getProducts(w http.ResponseWriter, r *http.Request){
	handleCors(w)

	if r.Method != "GET" {
		http.Error(w, "Please give me a GET request",400)
		return
	}
	sendData(w, productList, 200)
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
 }

 func createProducts(w http.ResponseWriter, r *http.Request){
	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != "POST" {
		http.Error(w, "Please give me a POST request",400)
		return
	}
	// r.body => description, title, price, imageUrl, id=> product er ekta
	// instance banabo => product list => append
	/*
	1. take body information (description, title, price, imageUrl, id) from r.Body
	2. create an instance using product struct with the body information
	3. append the instance into the product list
	*/

	var newProduct Product 
	decoder :=json.NewDecoder(r.Body)
	err :=decoder.Decode(&newProduct)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "Please give me a valid json",400)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct,201)
 }


 func handleCors (w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE,     OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
 }

 func handlePreflightReq (w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
 }
 
 func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
 }

func main() {
	mux:= http.NewServeMux()  // mux is a router

	mux.HandleFunc("/", helloHandler)  // route

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc ("/products", getProducts)

	mux.HandleFunc("/create-products", createProducts)

	fmt.Println("Server started on: 3000")
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)	
	}
}

func init () {
	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Description 1",
		Price: 100.00,
		ImgURL: "https://static.wikia.nocookie.net/colors/images/4/42/Orange_%28fruit%29.png/revision/latest/scale-to-width-down/985?cb=20250701034836",
	}
	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Description 2",
		Price: 200.00,
		ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSsT3iUXGmyhhGe8I6l5RAu1SEoz6ybubyKKQ&s",
	}
	prd3 := Product{
		ID: 3,
		Title: "Mango",
		Description: "Description 3",
		Price: 300.00,
		ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}
	prd4 := Product{
		ID: 4,
		Title: "Banana",
		Description: "Description 4",
		Price: 400.00,
		ImgURL: "https://dtgxwmigmg3gc.cloudfront.net/imagery/assets/derivations/icon/512/512/true/eyJpZCI6IjIzNjFhNjAyM2YxMmY1OGM5NzcyNTgwNjM3YWZiNjAxLmpwZyIsInN0b3JhZ2UiOiJwdWJsaWNfc3RvcmUifQ?signature=7d201f5f9b7458fa9243e619222f7453c4266c06112ca287d01adaf32bdb2d1f",
	}
	prd5 := Product{
		ID: 5,
		Title: "Pineapple",
		Description: "Description 5",
		Price: 500.00,
		ImgURL: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/organic%20pineaple.png?itok=4wB5t6Tk-24IDL_Zc",
	}
	
	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
	
}