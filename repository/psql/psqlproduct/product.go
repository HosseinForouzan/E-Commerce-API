package psqlproduct

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
)

func (d *DB) GetProducts(p param.ProductRequest) ([] entity.Product, error){
	var product entity.Product
	var products []entity.Product
	query := "SELECT * FROM products WHERE 1=1 "
	if p.CategoryID != 0 {
		query += fmt.Sprintf("AND category_id = %d ", p.CategoryID)
	}

	if p.Search != "" {
		query += fmt.Sprintf("AND name='%s' ", p.Search)
	}

	rows, err := d.conn.Conn().Query(context.Background(), query)
	if err != nil {
		return [] entity.Product{}, fmt.Errorf("Can't get products: %w", err)
	}


	for rows.Next() {
			rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
			products = append(products, product)
	}

	return products, nil
}