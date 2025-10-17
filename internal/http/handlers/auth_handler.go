package handlers

import (
	"net/http"

	"belajar-go/internal/core/ports"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ svc ports.AuthService }

func NewAuthHandler(svc ports.AuthService) *AuthHandler { return &AuthHandler{svc: svc} }

func (h *AuthHandler) Signup(c *gin.Context) {
	var in struct{ Email, Password string }
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	u, err := h.svc.Signup(c, in.Email, in.Password)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, u)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var in struct{ Email, Password string }
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	token, u, err := h.svc.Login(c, in.Email, in.Password)
	if err != nil {
		resp.Unauthorized(c, err.Error())
		return
	}
	// set cookie optional
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "/", "", false, true)
	resp.OK(c, gin.H{"token": token, "user": u})
}
