package postgres

import (
	"ganhaum.henrybarreto.dev/pkg/models"
)

func (r *Store) CreateProduct(product *models.ProductData) (*models.Product, error) {
	var model models.Product

	rows, err := r.db.Query("INSERT INTO products (thumbnail, name, description, type, value) VALUES ($1, $2, $3, $4, $5) RETURNING *", product.Thumbnail, product.Name, product.Description, product.Type, product.Value)
	if err != nil {
		return nil, NewErrProductQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Thumbnail, &model.Name, &model.Description, &model.Type, &model.Value, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			return nil, NewErrProductScan(err)
		}
	}

	return &model, nil
}

func (r *Store) GetProduct(id int) (*models.Product, error) {
	var model models.Product

	rows, err := r.db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return nil, NewErrProductQuery(err)
	}

	if !rows.NextResultSet() {
		return nil, NewErrProductNotFound(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Thumbnail, &model.Name, &model.Description, &model.Type, &model.Value, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			return nil, NewErrProductScan(err)
		}
	}

	return &model, nil
}

func (r *Store) GetProducts(page int, limit int) ([]models.Product, error) {
	rows, err := r.db.Query("SELECT * FROM products LIMIT $1 OFFSET $2", limit, (page-1)*limit)
	if err != nil {
		return nil, NewErrProductQuery(err)
	}

	if !rows.NextResultSet() {
		return nil, NewErrProductsEmpty(err)
	}

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.ID, &p.Thumbnail, &p.Name, &p.Description, &p.Type, &p.Value, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, NewErrProductsScan(err)
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *Store) UpdateProduct(id int, product *models.ProductData) (*models.Product, error) {
	var model models.Product

	rows, err := r.db.Query("UPDATE products SET thumbnail = $1, name = $2, description = $3, type = $4, value = $5 WHERE id = $6 RETURNING *", product.Thumbnail, product.Name, product.Description, product.Type, product.Value, id)
	if err != nil {
		return nil, NewErrProductQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Thumbnail, &model.Name, &model.Description, &model.Type, &model.Value, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			return nil, NewErrProductScan(err)
		}
	}

	return &model, nil
}

func (r *Store) DeleteProduct(id int) error {
	_, err := r.db.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return NewErrProductQuery(err)
	}

	return nil
}
