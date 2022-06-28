package main

import (
	"net/http"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

var items []Item

func main() {
	router := http.NewServeMux()
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`"{"Message": "sukses"}"`))
	// })

	router.HandleFunc("/get-item", GetItem)
	router.HandleFunc("/update-item", UpdateItem)
	router.HandleFunc("/create-item", CreateItem)
	router.HandleFunc("/delete-item", DeleteItem)

	http.ListenAndServe(":4000", router)
}
