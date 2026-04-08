package models

type Stock struct {
	ID              int         `json:"id"`
	Description     string      `json:"description"`
	Location        string      `json:"location"`
	MinimumQuantity map[int]int `json:"minimumQuantity"`
}

func newStock(description string, location string, minimumQuantity map[int]int) *Stock {

	return &Stock{Description: description, Location: location, MinimumQuantity: minimumQuantity}
}
