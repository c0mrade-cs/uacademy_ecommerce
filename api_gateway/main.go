package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/uacademy/e_commerce/api_gateway/clients"
	"github.com/uacademy/e_commerce/api_gateway/config"
	"github.com/uacademy/e_commerce/api_gateway/docs"
	"github.com/uacademy/e_commerce/api_gateway/handlers"
)

func main() {
	cfg := config.Load()

	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	//template GET method
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	h := handlers.Handler{
		GrpcClients: grpcClients,
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/order", h.CreateOrder)
		v1.GET("/order", h.GetOrderList)
		v1.GET("/order/:id", h.GetOrderById)

		v1.POST("/category", h.CreateCategory)
		v1.GET("/category/:id", h.GetCategoryById)
		v1.GET("/category", h.GetCategoryList)
		v1.PUT("/category", h.UpdateCategory)
		v1.DELETE("/category/:id", h.DeleteCategory)

		v1.POST("/product", h.CreateProduct)
		v1.GET("/product/:id", h.GetProductById)
		v1.GET("/product", h.GetProductList)
		v1.PUT("/product", h.UpdateProduct)
		v1.DELETE("/product/:id", h.DeleteProduct)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(cfg.HTTPPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
