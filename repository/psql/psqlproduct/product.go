package psqlproduct

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/jackc/pgx/v5"
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

func (d *DB) GetProductByID(id uint) (entity.Product, error) {
	var product entity.Product
	err := d.conn.Conn().QueryRow(context.Background(), "SELECT * FROM products WHERE id=$1", id).
	Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return entity.Product{}, fmt.Errorf("Can't get product by id: %w", err)
	}

	return product, nil
}

func (d *DB) CreateProduct(p entity.Product) (entity.Product, error) {
	var id uint

	err := d.conn.Conn().QueryRow(context.Background(),
									`INSERT INTO products(name,description,price,stock,category_id) 
									 VALUES($1,$2,$3,$4,$5) RETURNING id`, p.Name, p.Description, p.Price, p.Stock, p.CategoryID).Scan(&id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("Can't insert product: %w", err)
	}	

	
	p.ID = id

	return p, nil

}

func (d *DB) CreateCategory(c entity.Category) (entity.Category, error) {
	var id uint

	err := d.conn.Conn().QueryRow(context.Background(), 
									`INSERT INTO categories(name) VALUES($1) RETURNING id`, c.Name).Scan(&id)
	if err != nil {
		return entity.Category{}, fmt.Errorf("can't insert category %w", err)
	}							

	c.ID = id
	
	return c, err
}

func (d *DB) UpdateProduct(p entity.Product) (entity.Product, error) {
	fmt.Println(p)
	_, err := d.conn.Conn().Exec(context.Background(),
									`UPDATE products SET name=$1, description=$2, price=$3, stock=$4,category_id=$5 WHERE id=$6`,
									p.Name, p.Description, p.Price, p.Stock, p.CategoryID, p.ID)
	if err != nil {
		return entity.Product{}, fmt.Errorf("Can't update product: %w", err)
	}

	return p, nil
}

func (d * DB) DeleteProduct(id uint) error {
	_, err := d.conn.Conn().Exec(context.Background(), "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("can't delete product: %w", err)
	}

	return nil
}

func (d *DB) DecreaseStockTx(ctx context.Context, tx pgx.Tx, productID uint, quantity uint) error {
	cmd, err := tx.Exec(
		ctx,
		`
		UPDATE products
		SET stock = stock - $1,
		    updated_at = NOW()
		WHERE id = $2
		  AND stock >= $1
		`,
		quantity,
		productID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("not enough stock or product not found")
	}

	return nil
}
