package common

type Order struct {
	ID         int     `json:"id"`
	Product    string  `json:"product"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	CustomerID int     `json:"customer_id"`
}
