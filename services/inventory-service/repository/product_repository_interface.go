package repository

import (
	"context"
	"inventory/models"
)

// InventoryRepository define as operações que nosso banco deve suportar
type ProductRepository interface {
	InsertProduct(ctx context.Context, product *models.Product) (int, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	GetProduct(ctx context.Context, productID int) (*models.Product, error)
	DeleteProduct(ctx context.Context, productID int) (int, error)
	ReduceStock(ctx context.Context, productID int) error
}
