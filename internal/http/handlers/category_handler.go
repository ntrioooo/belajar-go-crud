package handlers

import (
	"strconv"

	"belajar-go/internal/core/ports"
	"belajar-go/internal/http/dto"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{ svc ports.CategoryService }

func NewCategoryHandler(svc ports.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var in struct {
		Name string `json:"name" binding:"required,min=3"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	cat, err := h.svc.Create(c.Request.Context(), in.Name)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.Created(c, cat)
}

func (h *CategoryHandler) List(c *gin.Context) {
	cats, err := h.svc.List(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}

	out := make([]dto.CategoryDTO, 0, len(cats))
	for i := range cats {
		out = append(out, *dto.NewCategoryDTO(&cats[i].Category))
	}
	resp.OK(c, out)
}

func (h *CategoryHandler) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil || cat == nil {
		resp.NotFound(c, "category not found")
		return
	}
	resp.OK(c, cat)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in struct{ Name string }
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	cat, err := h.svc.Update(c, uint(id), in.Name)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, cat)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.Delete(c, uint(id)); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"message": "Successfully to deleted category", "status": "success"})
}
