package handlers

import (
	"strconv"

	"belajar-go/internal/core/ports"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type PostHandler struct{ svc ports.PostService }

func NewPostHandler(svc ports.PostService) *PostHandler { return &PostHandler{svc: svc} }

func (h *PostHandler) Create(c *gin.Context) {
	var in struct{ Title, Body string }
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	// contoh: ambil user dari context (diisi oleh auth middleware)
	// user := c.MustGet("user").(domain.User)
	p, err := h.svc.Create(c, 0, in.Title, in.Body)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.Created(c, p)
}

func (h *PostHandler) List(c *gin.Context) {
	ps, err := h.svc.List(c)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, ps)
}

func (h *PostHandler) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.svc.GetByID(c, uint(id))
	if err != nil || p == nil {
		resp.NotFound(c, "post not found")
		return
	}
	resp.OK(c, p)
}

func (h *PostHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in struct{ Title, Body string }
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	p, err := h.svc.Update(c, uint(id), in.Title, in.Body)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, p)
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.Delete(c, uint(id)); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.NoContent(c)
}
