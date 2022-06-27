package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Qty   string `json:"qty"`
}

var items []Item

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`"{"Message": "sukses"}"`))
	})

	router.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(items)
		case "POST":
			var item Item
			json.NewDecoder(r.Body).Decode(&item)
			items = append(items, item)
			json.NewEncoder(w).Encode(item)
		case "PUT":
			query := r.URL.Query()
			id, _ := strconv.Atoi(query.Get("id"))
			for index, item := range items {
				json.NewDecoder(r.Body).Decode(&item)
				if item.ID == id {
					items[index].ID = item.ID
					items[index].Name = item.Name
					w.Write([]byte("Success to update item"))
				}
			}
		case "DELETE":
			query := r.URL.Query()
			id, _ := strconv.Atoi(query.Get("id"))
			for index, item := range items {
				if item.ID == id {
					items = append(items[:index], items[index+1:]...)
					w.Write([]byte("Success to delete item"))
				}
			}
		}
	})

	http.ListenAndServe(":4000", router)
}
