package postgres

import (
	"errors"
	"time"

	e "github.com/uacademy/e_commerce/catalog_service/proto-gen/ecommerce"
)

func (stg Postgres) CreateProduct(id string, entity *e.CreateProductRequest) error {
	_, err := stg.GetCategoryById(entity.CategoryId)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO product 
	(
		id, 
		category_id, 
		title, 
		descript, 
		price
	) VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5
	)`,
		id,
		entity.CategoryId,
		entity.Title,
		entity.Description,
		entity.Price,
	)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) GetProductList(offset, limit int, search string) (resp *e.GetProductListResponse, err error) {
	resp = &e.GetProductListResponse{
		Products: make([]*e.Product, 0),
	}

	rows, err := stg.db.Queryx(`SELECT
	id,
	title,
	descript,
	price,
	created_at,
	updated_at
	FROM product WHERE deleted_at IS NULL AND (title ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		p := &e.Product{}
		var updatedAt *string

		err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil {
			p.UpdatedAt = *updatedAt
		}

		resp.Products = append(resp.Products, p)
	}

	return resp, err
}

func (stg Postgres) GetProductById(id string) (*e.GetProductByIdResponse, error) {
	res := &e.GetProductByIdResponse{
		Category: &e.GetProductByIdResponse_Category{},
	}

	var deletedAt *time.Time
	var updatedAt, categoryUpdatedAt *string

	err := stg.db.QueryRow(`SELECT 
		pr.id, 
		pr.title, 
		pr.descript, 
		pr.price, 
		pr.created_at, 
		pr.updated_at, 
		pr.deleted_at,
		ca.id, 
		ca.category_title, 
		ca.created_at, 
		ca.updated_at 
	FROM product AS pr JOIN category AS ca ON pr.category_id = ca.id WHERE pr.id=$1`, id).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Price,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
		&res.Category.Id,
		&res.Category.CategoryTitle,
		&res.Category.CreatedAt,
		&categoryUpdatedAt,
	)
	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if categoryUpdatedAt != nil {
		res.Category.UpdatedAt = *categoryUpdatedAt
	}

	if deletedAt != nil {
		return res, errors.New("product not found")
	}

	return res, nil
}

func (stg Postgres) UpdateProduct(entity *e.UpdateProductRequest) error {
	res, err := stg.db.NamedExec(`UPDATE product  SET title=:t, descript=:d, price=:p, updated_at=now() WHERE deleted_at IS NULL AND id=:id`, map[string]interface{}{
		"id": entity.Id,
		"t":  entity.Title,
		"d":  entity.Description,
		"p":  entity.Price,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("product not found")
}

func (stg Postgres) DeleteProduct(id string) error {
	res, err := stg.db.Exec(`UPDATE "product"  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}
	return errors.New("product not found")
}
