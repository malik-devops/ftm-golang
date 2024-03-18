package entity

//Product object for REST(CRUD)
type Product struct {
	ID        int    `json:"id"`
	Name string `json:"name"`
	Description  string `json:"description"`
	Quantity       int    `json:"quantity"`
}