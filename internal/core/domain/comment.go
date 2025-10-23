package domain

import "time"

// Comment merepresentasikan komentar pada sebuah post.
// Mendukung threaded replies via ParentID (nullable).
type Comment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	ParentID  *uint     `json:"parent_id,omitempty"` // null = komentar level-1
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentView untuk response: menambahkan informasi penulis,
// like count, apakah dilike oleh current user, dan jumlah balasan.
type CommentView struct {
	Comment
	AuthorUsername string `json:"author_username"`
	LikeCount      int64  `json:"like_count"`
	LikedByMe      bool   `json:"liked_by_me"`
	ReplyCount     int64  `json:"reply_count"`
}
