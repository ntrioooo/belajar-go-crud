package handlers

import (
	"strconv"

	"belajar-go/internal/core/ports"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type RetweetHandler struct {
	svc ports.RetweetService
}

func NewRetweetHandler(svc ports.RetweetService) *RetweetHandler {
	return &RetweetHandler{svc: svc}
}

// POST /v1/posts/:id/retweet
// Toggle pure retweet (tanpa quote)
func (h *RetweetHandler) ToggleRetweet(c *gin.Context) {
	postID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	userID := uidAny.(uint)

	retweeted, count, err := h.svc.Toggle(c, userID, uint(postID64), nil)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"retweeted": retweeted, "retweet_count": count})
}

// POST /v1/posts/:id/quote
// Toggle retweet dengan quote body (kalau sudah ada, akan unretweet)
func (h *RetweetHandler) ToggleQuote(c *gin.Context) {
	postID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	userID := uidAny.(uint)

	var in struct {
		QuoteBody *string `json:"quote_body"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	retweeted, count, err := h.svc.Toggle(c, userID, uint(postID64), in.QuoteBody)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"retweeted": retweeted, "retweet_count": count})
}
