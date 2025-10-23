package handlers

import (
	"strconv"

	"belajar-go/internal/core/ports"
	"belajar-go/internal/http/dto"
	"belajar-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	svc ports.CommentService
}

func NewCommentHandler(svc ports.CommentService) *CommentHandler {
	return &CommentHandler{svc: svc}
}

// POST /v1/posts/:id/comments
func (h *CommentHandler) Create(c *gin.Context) {
	postID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	userID := uidAny.(uint)

	var in struct {
		Body     string `json:"body" binding:"required"`
		ParentID *uint  `json:"parent_id"` // optional for reply
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	comment, err := h.svc.Create(c, userID, uint(postID64), in.ParentID, in.Body)
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.Created(c, comment)
}

// GET /v1/posts/:id/comments
// (ambil semua komentar untuk post; untuk reply detail gunakan endpoint replies)
func (h *CommentHandler) ListByPost(c *gin.Context) {
	postID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var viewer uint
	if v, ok := c.Get("userID"); ok {
		viewer = v.(uint)
	}

	views, err := h.svc.GetByPost(c, viewer, uint(postID64))
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, dto.NewCommentViewDTOs(views))
}

// GET /v1/comments/:id/replies
func (h *CommentHandler) ListReplies(c *gin.Context) {
	parentID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var viewer uint
	if v, ok := c.Get("userID"); ok {
		viewer = v.(uint)
	}

	views, err := h.svc.GetReplies(c, viewer, uint(parentID64))
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, views)
}

// POST /v1/comments/:id/like  (toggle like)
func (h *CommentHandler) ToggleLike(c *gin.Context) {
	commentID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	userID := uidAny.(uint)

	liked, count, err := h.svc.ToggleLike(c, userID, uint(commentID64))
	if err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"liked": liked, "like_count": count})
}

// DELETE /v1/comments/:id
func (h *CommentHandler) Delete(c *gin.Context) {
	commentID64, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	uidAny, ok := c.Get("userID")
	if !ok {
		resp.Unauthorized(c, "unauthorized")
		return
	}
	userID := uidAny.(uint)

	if err := h.svc.Delete(c, userID, uint(commentID64)); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"message": "Successfully deleted comment", "status": "success"})
}
