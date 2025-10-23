package domain

import "time"

// Retweet merepresentasikan aksi user "meneruskan" sebuah post.
// Unique(user_id, original_post_id) agar tidak dobel.
type Retweet struct {
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id"`
	OriginalPostID uint      `json:"original_post_id"`
	QuoteBody      *string   `json:"quote_body,omitempty"` // null jika pure retweet (tanpa komentar)
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Ringkasan post untuk ditampilkan bersama RetweetView
type PostSummary struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	AuthorUsername string `json:"author_username"`
	CategoryName   string `json:"category_name"`
}

// RetweetView: menampilkan retweet + ringkasan original post.
type RetweetView struct {
	Retweet
	Original PostSummary `json:"original"`
}
