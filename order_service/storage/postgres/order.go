package postgres

import (
	"errors"

	e "github.com/uacademy/e_commerce/order_service/proto-gen/ecommerce"
)

func (stg Postgres) CreateOrder(id string, entity *e.CreateOrderRequest) error {
	_, err := stg.db.Exec(`INSERT INTO "order" 
	(
		id, 
		product_id, 
		quantity, 
		customer_name, 
		customer_address, 
		customer_phone) 
	VALUES 
	(
		$1, 
		$2, 
		$3, 
		$4, 
		$5, 
		$6
	)`,
		id,
		entity.ProductId,
		entity.Quantity,
		entity.CustomerName,
		entity.CustomerAddress,
		entity.CustomerPhone,
	)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) GetOrderList(offset, limit int, search string) (resp *e.GetOrderListResponse, err error) {
	resp = &e.GetOrderListResponse{
		Orders: make([]*e.Order, 0),
	}

	rows, err := stg.db.Queryx(`SELECT
	id,
	product_id,
	quantity,
	customer_name,
	customer_address,
	customer_phone,
	created_at
	FROM "order" WHERE (customer_address ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		o := &e.Order{}

		err := rows.Scan(
			&o.Id,
			&o.ProductId,
			&o.Quantity,
			&o.CustomerName,
			&o.CustomerAddress,
			&o.CustomerPhone,
			&o.CreatedAt,
		)
		if err != nil {
			return resp, err
		}

		resp.Orders = append(resp.Orders, o)
	}

	return resp, err
}

func (stg Postgres) GetOrderById(id string) (*e.GetOrderByIdResponse, error) {
	res := &e.GetOrderByIdResponse{
		Product: &e.GetOrderByIdResponse_Product{},
	}

	err := stg.db.QueryRow(`SELECT
		id, 
		product_id, 
		quantity, 
		customer_name, 
		customer_address, 
		customer_phone, 
		created_at
	FROM "order" WHERE id = $1`, id).Scan(
		&res.Id,
		&res.Product.Id,
		&res.Quantity,
		&res.CustomerName,
		&res.CustomerAddress,
		&res.CustomerPhone,
		&res.CreatedAt,
	)
	if err != nil {
		return res, errors.New("order not found")
	}

	return res, nil
}
