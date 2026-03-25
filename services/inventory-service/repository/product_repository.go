package repository

import (
	"context"
	"errors"
	"inventory/database"
	"inventory/models"
)

type productRepo struct {
	db *database.DBInventory
}

func NewProductRepo(db *database.DBInventory) *productRepo {
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

func (r *productRepo) DeleteProduct(ctx context.Context, id int) (int, error) {
	// Subtraímos a quantidade do estoque atual
	query := "DELETE FROM product where id=$1"

	result, err := r.db.Conn.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rows, _ := result.RowsAffected()

	if rows == 0 {
		return 0, errors.New("Item does't exists")
	}

	return int(rows), nil
}

func (r *productRepo) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	query := "SELECT FROM product where id=$1"
	var product *models.Product
	r.db.Conn.QueryRowContext(ctx, query, id).Scan(&product)

	return product, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	return nil, nil
}
