package models

type Order struct {
	ProductID   int     `json:"product_id"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	Value       float32 `json:"value"`
}
