package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/uacademy/e_commerce/catalog_service/config"
	e "github.com/uacademy/e_commerce/catalog_service/proto-gen/ecommerce"
	catSrv "github.com/uacademy/e_commerce/catalog_service/services/category"
	prodSrv "github.com/uacademy/e_commerce/catalog_service/services/product"
	"github.com/uacademy/e_commerce/catalog_service/storage"
	"github.com/uacademy/e_commerce/catalog_service/storage/postgres"
)

func main() {
	cfg := config.Load()

	psqlConString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var stg storage.StorageI
	stg, err := postgres.InitDb(psqlConString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	e.RegisterCategoryServiceServer(s, catSrv.NewCategoryService(stg))
	e.RegisterProductServiceServer(s, prodSrv.NewProductService(stg))
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
