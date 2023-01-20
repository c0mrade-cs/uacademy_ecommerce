package clients

import (
	"github.com/uacademy/e_commerce/api_gateway/config"
	e "github.com/uacademy/e_commerce/api_gateway/proto-gen/ecommerce"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Category e.CategoryServiceClient
	Product  e.ProductServiceClient
	Order    e.OrderServiceClient

	conns []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connCategory, err := grpc.Dial(cfg.CatalogServiceGrpcHost+cfg.CatalogServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	category := e.NewCategoryServiceClient(connCategory)

	connProduct, err := grpc.Dial(cfg.CatalogServiceGrpcHost+cfg.CatalogServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	product := e.NewProductServiceClient(connProduct)

	connOrder, err := grpc.Dial(cfg.OrderServiceGrpcHost+cfg.OrderServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	order := e.NewOrderServiceClient(connOrder)

	conns := make([]*grpc.ClientConn, 0)

	return &GrpcClients{
		Category: category,
		Product:  product,
		Order:    order,
		conns:    append(conns, connCategory, connProduct, connOrder),
	}, nil
}
