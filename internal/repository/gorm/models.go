package gormrepo

import "time"

// Entity yang disimpan DB untuk GORM (bisa berbeda dari domain)
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Username  string `gorm:"uniqueIndex;size:50"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

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

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	PostID    uint `gorm:"index"`
	CreatedAt time.Time
	// UNIQUE (user_id, post_id)
}

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex;size:100"`
	UpdatedAt time.Time
	CreatedAt time.Time
	// UNIQUE (user_id, post_id)
}
