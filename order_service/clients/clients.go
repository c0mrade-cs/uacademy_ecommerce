package clients

import (
	"github.com/uacademy/e_commerce/order_service/config"
	e "github.com/uacademy/e_commerce/order_service/proto-gen/ecommerce"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Product  e.ProductServiceClient
	Category e.CategoryServiceClient
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

	return &GrpcClients{
		Category: category,
		Product:  product,
	}, nil
}
