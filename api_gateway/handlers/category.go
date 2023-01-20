package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	"github.com/uacademy/e_commerce/api_gateway/models"
	e "github.com/uacademy/e_commerce/api_gateway/proto-gen/ecommerce"
)

// CreateCategory godoc
// @Summary     Create category
// @Description create new category
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       category      body     models.CreateCategoryModel true  "Category body"
// @Success     201           {object} models.JSONResponse{data=models.Category}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     500           {object} models.JSONErrorResponse
// @Router      /v1/category [post]
func (h Handler) CreateCategory(c *gin.Context) {
	var body models.CreateCategoryModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	category, err := h.GrpcClients.Category.CreateCategory(c.Request.Context(), &e.CreateCategoryRequest{
		CategoryTitle: body.CategoryTitle,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Category Created",
		Data:    category,
	})
}

// GetCategory godoc
// @Summary     Get category
// @Description get category by ID
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Category ID"
// @Success     200           {object} models.JSONResponse{data=models.Category}
// @Failure     404           {object} models.JSONErrorResponse
// @Router      /v1/category/{id} [get]
func (h Handler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	category, err := h.GrpcClients.Category.GetCategoryById(c.Request.Context(), &e.GetCategoryByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    category,
	})
}

// ListCategorys godoc
// @Summary     List categories
// @Description get categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       offset        query    string false "0"
// @Param       limit         query    string false "10"
// @Param       search        query    string false "smth"
// @Success     200           {object} models.JSONResponse{data=[]models.Category}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     500           {object} models.JSONErrorResponse
// @Router      /v1/category [get]
func (h Handler) GetCategoryList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	categoryList, err := h.GrpcClients.Category.GetCategoryList(c.Request.Context(), &e.GetCategoryListRequest{
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
		Data:    categoryList,
	})
}

// UpdateCategory godoc
// @Summary     Update category
// @Description update category
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       category      body     models.UpdateCategoryModel true  "Category body"
// @Success     200           {object} models.JSONResponse{data=models.Category}
// @Failure     400           {object} models.JSONErrorResponse
// @Failure     404           {object} models.JSONErrorResponse
// @Router      /v1/category [put]
func (h Handler) UpdateCategory(c *gin.Context) {
	var body models.UpdateCategoryModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	category, err := h.GrpcClients.Category.UpdateCategory(c.Request.Context(), &e.UpdateCategoryRequest{
		Id:            body.Id,
		CategoryTitle: body.CategoryTitle,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Category Updated",
		Data:    category,
	})
}

// DeleteCategory godoc
// @Summary     Delete category
// @Description delete category by ID
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Category ID"
// @Success     200           {object} models.JSONResponse{data=models.Category}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/category/{id} [delete]
func (h Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.GrpcClients.Category.DeleteCategory(c.Request.Context(), &e.DeleteCategoryRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Category Deleted",
		Data:    category,
	})
}
