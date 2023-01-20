package order

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/uacademy/e_commerce/order_service/clients"
	e "github.com/uacademy/e_commerce/order_service/proto-gen/ecommerce"
	"github.com/uacademy/e_commerce/order_service/storage"

	"context"
)

type orderService struct {
	stg storage.StorageI
	e.UnimplementedOrderServiceServer
	GrpcClients *clients.GrpcClients
}

// NewOrderService ...
func NewOrderService(client *clients.GrpcClients, stg storage.StorageI) *orderService {
	return &orderService{
		stg:         stg,
		GrpcClients: client,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, req *e.CreateOrderRequest) (*e.Order, error) {
	id := uuid.New()

	err := s.stg.CreateOrder(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.CreateOrder: %s", err.Error())
	}

	order, err := s.stg.GetOrderById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetOrderByID: %s", err.Error())
	}

	return &e.Order{
		Id:              id.String(),
		ProductId:       order.Product.Id,
		Quantity:        order.Quantity,
		CustomerName:    order.CustomerName,
		CustomerAddress: order.CustomerAddress,
		CustomerPhone:   order.CustomerPhone,
		CreatedAt:       order.CreatedAt,
	}, nil
}

func (s *orderService) GetOrderById(ctx context.Context, req *e.GetOrderByIdRequest) (*e.GetOrderByIdResponse, error) {
	order, err := s.stg.GetOrderById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "s.stg.GetOrderById: %s", err.Error())
	}

	return order, nil
}

func (s *orderService) GetOrderList(ctx context.Context, req *e.GetOrderListRequest) (*e.GetOrderListResponse, error) {
	res, err := s.stg.GetOrderList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetOrderList: %s", err.Error())
	}

	return res, nil
}
