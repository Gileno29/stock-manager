package models

import "errors"

type Product struct {
	ID            int     `json:"id"`
	Code          int     `json:"code"`
	Quantity      int     `json:"quantity"`
	Description   string  `json:"description"`
	Value         float32 `json:"value"`
	StockLocation Stock
}

func (p *Product) AddToStock(s *Stock) error {
	if s == nil {
		return errors.New("Error to add in stock, stock don't exists")
	}
	s.MinimumQuantity[p.Code]++
	p.StockLocation = *s
	return nil
}

func (p *Product) RemoveFromStrock() (error, string) {

	if p.StockLocation.MinimumQuantity[p.Code] == 0 {
		return errors.New("You dont have products in your stock"), ""
	}
	p.StockLocation.MinimumQuantity[p.Code]--

	if p.StockLocation.MinimumQuantity[p.Code] < 2 {
		return nil, "You have few products in your stock"
	}

	return nil, "Product Successfull removed"
}
