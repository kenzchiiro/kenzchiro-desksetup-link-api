package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenzchiro/desksetup-link-api/domain"
	"github.com/kenzchiro/desksetup-link-api/services/product"
)

type ProductHandler struct {
	service *product.ProductService
}

func NewProductHandler(service *product.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/products")
	{
		api.GET("/", h.GetAll)
		api.GET("/:id", h.GetByID)
		api.POST("/", h.Create)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: products, Count: len(products)})
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid product id"})
		return
	}

	product, found, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "product not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: product})
}

func (h *ProductHandler) Create(c *gin.Context) {
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid request body"})
		return
	}

	created, err := h.service.Create(c.Request.Context(), p)
	if err != nil {
		status := http.StatusInternalServerError
		if err == domain.ErrInvalidProductTitle {
			status = http.StatusBadRequest
		}
		c.JSON(status, Response{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Success: true, Message: "Product created successfully", Data: created})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid product id"})
		return
	}

	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid request body"})
		return
	}

	updated, found, err := h.service.Update(c.Request.Context(), id, p)
	if err != nil {
		status := http.StatusInternalServerError
		if err == domain.ErrInvalidProductTitle {
			status = http.StatusBadRequest
		}
		c.JSON(status, Response{Success: false, Error: err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "product not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "Product updated successfully", Data: updated})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid product id"})
		return
	}

	deleted, err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "product not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "Product deleted successfully"})
}

func parseID(raw string) (int64, error) {
	return strconv.ParseInt(raw, 10, 64)
}
