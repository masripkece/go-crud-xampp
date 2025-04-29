package main

import (
	"fmt"
	"go-crud-xampp/config"
	"go-crud-xampp/handlers"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/kategori/create", handlers.CreateKategori)
	http.HandleFunc("/produk/create", handlers.CreateProduk)
	http.HandleFunc("/produk/read", handlers.GetProduk)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
