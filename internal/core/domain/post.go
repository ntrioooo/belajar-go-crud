package domain

import "time"

type Post struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	UserID     uint      `json:"user_id"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// View model untuk list/show
type PostView struct {
	Post
	AuthorUsername string `json:"author_username"`
	CategoryName   string `json:"category_name"`
	LikeCount      int64  `json:"like_count"`
	LikedByMe      bool   `json:"liked_by_me"`

	// Tambahan
	CommentCount int64 `json:"comment_count"`
	RetweetCount int64 `json:"retweet_count"`
	// Jika item di timeline adalah "retweet event", bisa set RepostedByUsername
	// dan/atau QuoteBody agar frontend tahu ini tampil sebagai retweet/quote.
	RepostedByUsername *string `json:"reposted_by_username,omitempty"`
	QuoteBody          *string `json:"quote_body,omitempty"`
}
