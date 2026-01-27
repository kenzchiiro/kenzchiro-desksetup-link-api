package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenzchiro/desksetup-link-api/domain"
	"github.com/kenzchiro/desksetup-link-api/service"
)

type HighlightHandler struct {
	service *service.HighlightService
}

func NewHighlightHandler(svc *service.HighlightService) *HighlightHandler {
	return &HighlightHandler{service: svc}
}

func (h *HighlightHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/highlights")
	{
		api.GET("", h.GetAll)
		api.GET("/", h.GetAll)
		api.GET("/:id", h.GetByID)
		api.POST("", h.Create)
		api.POST("/", h.Create)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}

func (h *HighlightHandler) GetAll(c *gin.Context) {
	highlights, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: highlights, Count: len(highlights)})
}

func (h *HighlightHandler) GetByID(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid highlight id"})
		return
	}

	highlight, found, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "highlight not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: highlight})
}

func (h *HighlightHandler) Create(c *gin.Context) {
	var hl domain.Highlight
	if err := c.ShouldBindJSON(&hl); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid request body"})
		return
	}

	created, err := h.service.Create(c.Request.Context(), hl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Success: true, Message: "Highlight created successfully", Data: created})
}

func (h *HighlightHandler) Update(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid highlight id"})
		return
	}

	var hl domain.Highlight
	if err := c.ShouldBindJSON(&hl); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid request body"})
		return
	}

	updated, found, err := h.service.Update(c.Request.Context(), id, hl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "highlight not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "Highlight updated successfully", Data: updated})
}

func (h *HighlightHandler) Delete(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "invalid highlight id"})
		return
	}

	deleted, err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "highlight not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "Highlight deleted successfully"})
}
