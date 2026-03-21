package repository

import (
	"context"
	"inventory/database"
	"inventory/models"
)

type productRepo struct {
	db *database.DBInventory
}

func NewProductRepo(db *database.DBInventory) productRepo {
	return &productRepo{db: db}
}

func (r *productRepo) InsertProduct(ctx context.Context, p *models.Product) (int, error) {
	// Subtraímos a quantidade do estoque atual
	query := "INSERT into product (code,quantity, description, value, stockLocation) VALUES ($1,$2,$3,$4,$5)"

	result, err := r.db.Conn.ExecContext(ctx, query, p.Code, p.Quantity, p.Description, p.Value, p.StockLocation)
	if err != nil {
		return 0, err
	}

	rows, _ := result.RowsAffected()

	return int(rows), nil
}

func (r *productRepo) DeleteProduct(ctx context.Context, id int) error {
	// Subtraímos a quantidade do estoque atual
	query := "DELETE FROM product where id=1$"

	result, err := r.db.Conn.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rows, _ := result.RowsAffected()

	return int(rows), nil
}
