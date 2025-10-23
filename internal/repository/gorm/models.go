package gormrepo

import "time"

// ===================== USERS =====================

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Username  string `gorm:"uniqueIndex;size:50"`
	Password  string
	Role      string `gorm:"index;size:20;default:member"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// ===================== POSTS =====================

type Post struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Body       string
	UserID     uint `gorm:"index"`
	CategoryID uint `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

// ===================== LIKES (POST) =====================

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index;uniqueIndex:uniq_like_user_post"`
	PostID    uint `gorm:"index;uniqueIndex:uniq_like_user_post"`
	CreatedAt time.Time
	// UNIQUE (user_id, post_id)
}

// ===================== CATEGORIES =====================

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex;size:100"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

// ===================== COMMENTS =====================

type Comment struct {
	ID        uint  `gorm:"primaryKey"`
	PostID    uint  `gorm:"index;constraint:OnDelete:CASCADE"`
	UserID    uint  `gorm:"index;constraint:OnDelete:CASCADE"`
	ParentID  *uint `gorm:"index"` // null = top-level
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type CommentLike struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index;uniqueIndex:uniq_like_user_comment"`
	CommentID uint `gorm:"index;uniqueIndex:uniq_like_user_comment"`
	CreatedAt time.Time
	// UNIQUE (user_id, comment_id)
}

// ===================== RETWEETS / QUOTES =====================

type Retweet struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint `gorm:"index;uniqueIndex:uniq_user_original"`
	OriginalPostID uint `gorm:"index;uniqueIndex:uniq_user_original;constraint:OnDelete:CASCADE"`
	QuoteBody      *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
	// UNIQUE (user_id, original_post_id)
}
