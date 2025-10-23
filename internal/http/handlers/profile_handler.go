package handlers

import (
	"belajar-go/internal/core/ports"
	"belajar-go/internal/http/dto"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	svc   ports.ProfileService
	users ports.UserRepository // untuk resolve username -> id
}

func NewProfileHandler(svc ports.ProfileService, users ports.UserRepository) *ProfileHandler {
	return &ProfileHandler{svc: svc, users: users}
}

// GET /v1/profiles/:username
func (h *ProfileHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")

	u, err := h.users.FindByUsername(c, username)
	if err != nil || u == nil {
		resp.NotFound(c, "user not found")
		return
	}

	var viewer uint
	if v, ok := c.Get("userID"); ok {
		viewer = v.(uint)
	}

	view, err := h.svc.GetProfile(c, viewer, u.ID)
	if err != nil || view == nil {
		resp.BadRequest(c, "cannot build profile")
		return
	}
	resp.OK(c, dto.NewProfileDTO(view))
}
