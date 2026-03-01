package repository

import (
	"context"
	"fmt"
	"inventory/database"
)

type repositoryRepo struct {
	db *database.DBInventory
}

func NewRepo(db *database.DBInventory) InventoryRepository {
	return &repositoryRepo{db: db}
}

func (r *repositoryRepo) UpdateStock(ctx context.Context, productID int, quantity int) error {
	// Subtraímos a quantidade do estoque atual
	query := "UPDATE stock SET stock = stock - $1 WHERE id = $2 AND stock >= $1"

	result, err := r.db.Conn.ExecContext(ctx, query, quantity, productID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("estoque insuficiente ou produto não encontrado")
	}

	return nil
}

func (r *repositoryRepo) GetStock(ctx context.Context, productID int) (int, error) {
	var stock int
	query := "SELECT stock FROM stock WHERE id = $1"
	err := r.db.Conn.QueryRowContext(ctx, query, productID).Scan(&stock)
	return stock, err
}

func (r *repositoryRepo) InsertStock(ctx context.Context, productID int, quantity int) (int, error) {

	var stock int

	query := "INSERT INTO stock(product_id, quantity) VALUES($1, $2) RETURNING id"

	err := r.db.Conn.QueryRowContext(ctx, query, productID, quantity).Scan(&stock)

	return stock, err
}
