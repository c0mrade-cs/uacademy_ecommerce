package models

import "time"

type Order struct {
	Id               string    `json:"id"`
	Product_id       string    `json:"product_id"`
	Quantity         int32     `json:"quantity"`
	Customer_name    string    `json:"customer_name"`
	Customer_address string    `json:"user_address"`
	Customer_phone   string    `json:"customer_phone"`
	Created_at       time.Time `json:"created_at"`
}

type CreateOrderModel struct {
	Product_id       string `json:"product_id"`
	Quantity         int32  `json:"quantity"`
	Customer_name    string `json:"customer_name"`
	Customer_address string `json:"user_address"`
	Customer_phone   string `json:"customer_phone"`
}

type PackedOrderModel struct {
	Id               string    `json:"id"`
	Product_id       string    `json:"product_id"`
	Quantity         int32     `json:"quantity"`
	Customer_name    string    `json:"customer_name"`
	Customer_address string    `json:"user_address"`
	Customer_phone   string    `json:"customer_phone"`
	Created_at       time.Time `json:"created_at"`
	Product          Product   `json:"product"`
}
