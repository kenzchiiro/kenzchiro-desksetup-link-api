package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenzchiro/desksetup-link-api/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(svc *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: svc}
}

func (h *CategoryHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/categories", h.GetAll)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Success: true, Data: categories, Count: len(categories)})
}
