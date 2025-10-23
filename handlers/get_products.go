package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	
	util.SendData(w, database.ProductList, 200)
	encoder := json.NewEncoder(w)
	encoder.Encode(database.ProductList)
 }
