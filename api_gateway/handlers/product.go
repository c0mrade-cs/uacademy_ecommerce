package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	"github.com/uacademy/e_commerce/api_gateway/models"
	e "github.com/uacademy/e_commerce/api_gateway/proto-gen/ecommerce"
)

// CreateProduct godoc
// @Summary     Create product
// @Description create new product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product       body     models.CreateProductModel true  "Product body"
// @Success     201           {object} models.JSONResponse{data=models.Product}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     500           {object} models.JSONErrorResponse
// @Router      /v1/product [post]
func (h Handler) CreateProduct(c *gin.Context) {
	var body models.CreateProductModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	product, err := h.GrpcClients.Product.CreateProduct(c.Request.Context(), &e.CreateProductRequest{
		CategoryId:  body.CategoryId,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Product Created",
		Data:    product,
	})
}

// GetProduct godoc
// @Summary     Get product
// @Description get product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Product ID"
// @Success     200           {object} models.JSONResponse{data=models.Product}
// @Failure     404           {object} models.JSONErrorResponse
// @Router      /v1/product/{id} [get]
func (h Handler) GetProductById(c *gin.Context) {
	id := c.Param("id")

	product, err := h.GrpcClients.Product.GetProductById(c.Request.Context(), &e.GetProductByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    product,
	})
}

// ListProducts godoc
// @Summary     List products
// @Description get products
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       offset        query    string false "0"
// @Param       limit         query    string false "10"
// @Param       search        query    string false "smth"
// @Success     200           {object} models.JSONResponse{data=[]models.Product}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     500           {object} models.JSONErrorResponse
// @Router      /v1/product [get]
func (h Handler) GetProductList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	productList, err := h.GrpcClients.Product.GetProductList(c.Request.Context(), &e.GetProductListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    productList,
	})
}

// UpdateProduct godoc
// @Summary     Update product
// @Description update product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product       body     models.UpdateProductModel true  "Product body"
// @Success     200           {object} models.JSONResponse{data=models.Product}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     404           {object} models.JSONErrorResponse
// @Router      /v1/product [put]
func (h Handler) UpdateProduct(c *gin.Context) {
	var body models.UpdateProductModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	product, err := h.GrpcClients.Product.UpdateProduct(c.Request.Context(), &e.UpdateProductRequest{
		Id:          body.Id,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Product Updated",
		Data:    product,
	})
}

// DeleteProduct godoc
// @Summary     Delete product
// @Description delete product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Product ID"
// @Success     200           {object} models.JSONResponse{data=models.Product}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/product/{id} [delete]
func (h Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.GrpcClients.Product.DeleteProduct(c.Request.Context(), &e.DeleteProductRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Product Deleted",
		Data:    product,
	})
}
