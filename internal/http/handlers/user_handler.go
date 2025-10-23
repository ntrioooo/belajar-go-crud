package handlers

import (
	"belajar-go/internal/core/ports"
	"belajar-go/internal/http/dto"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc ports.UserService
}

func NewUserHandler(svc ports.UserService) *UserHandler { return &UserHandler{svc: svc} }

// GET /v1/users/me
func (h *UserHandler) Me(c *gin.Context) {
	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	user, err := h.svc.GetMe(c.Request.Context(), uidAny.(uint))
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, user)
}

// PUT /v1/users/me
func (h *UserHandler) UpdateMe(c *gin.Context) {
	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	var in struct {
		Email    *string `json:"email"`    // optional
		Username *string `json:"username"` // optional
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	user, err := h.svc.UpdateMe(c.Request.Context(), uidAny.(uint), in.Email, in.Username)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, dto.NewUserDTO(user))
}
