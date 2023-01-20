package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/uacademy/e_commerce/api_gateway/models"
	e "github.com/uacademy/e_commerce/api_gateway/proto-gen/ecommerce"
)

// CreateOrder godoc
// @Summary     Create order
// @Description create new order
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       order         body     models.CreateOrderModel true  "Order body"
// @Success     201           {object} models.JSONResponse{data=models.Order}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     500           {object} models.JSONErrorResponse
// @Router      /v1/order [post]
func (h Handler) CreateOrder(c *gin.Context) {
	var body models.CreateOrderModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	_, err := h.GrpcClients.Product.GetProductById(c.Request.Context(), &e.GetProductByIdRequest{
		Id: body.Product_id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	order, err := h.GrpcClients.Order.CreateOrder(c.Request.Context(), &e.CreateOrderRequest{
		ProductId:       body.Product_id,
		Quantity:        body.Quantity,
		CustomerName:    body.Customer_name,
		CustomerAddress: body.Customer_address,
		CustomerPhone:   body.Customer_phone,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Order Created",
		Data:    order,
	})
}

// GetOrder godoc
// @Summary     Get order
// @Description get order by ID
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Order ID"
// @Success     200           {object} models.JSONResponse{data=models.PackedOrderModel}
// @Failure     404           {object} models.JSONErrorResponse
// @Router      /v1/order/{id} [get]
func (h Handler) GetOrderById(c *gin.Context) {
	id := c.Param("id")

	order, err := h.GrpcClients.Order.GetOrderById(c.Request.Context(), &e.GetOrderByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	product, err := h.GrpcClients.Product.GetProductById(c.Request.Context(), &e.GetProductByIdRequest{
		Id: order.Product.Id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	order.Product = &e.GetOrderByIdResponse_Product{
		Id:    product.Id,
		Title: product.Title,
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    order,
	})
}

// ListOrders godoc
// @Summary     List orders
// @Description get orders
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       offset        query    string false "0"
// @Param       limit         query    string false "10"
// @Param       search        query    string false "smth"
// @Success     200           {object} models.JSONResponse{data=[]models.Order}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     500           {object} models.JSONErrorResponse
// @Router      /v1/order [get]
func (h Handler) GetOrderList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

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

	orderList, err := h.GrpcClients.Order.GetOrderList(c.Request.Context(), &e.GetOrderListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    orderList,
	})
}
