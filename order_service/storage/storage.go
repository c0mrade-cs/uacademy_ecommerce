package storage

import e "github.com/uacademy/e_commerce/order_service/proto-gen/ecommerce"

type StorageI interface {
	CreateOrder(id string, entity *e.CreateOrderRequest) error
	GetOrderList(offset, limit int, search string) (resp *e.GetOrderListResponse, err error)
	GetOrderById(id string) (*e.GetOrderByIdResponse, error)
}
