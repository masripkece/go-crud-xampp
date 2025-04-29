package models

type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
	Kategori   string `json:"kategori,omitempty"`
}
