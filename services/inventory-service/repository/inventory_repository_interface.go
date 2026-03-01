package repository

import "context"

// InventoryRepository define as operações que nosso banco deve suportar
type InventoryRepository interface {
	InsertStock(ctx context.Context, productID int, quantity int) (int, error)
	UpdateStock(ctx context.Context, productID int, quantity int) error
	GetStock(ctx context.Context, productID int) (int, error)
}
