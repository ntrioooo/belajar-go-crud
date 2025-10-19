package gormrepo

import "time"

// Entity yang disimpan DB untuk GORM (bisa berbeda dari domain)
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Body      string
	UserID    uint `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
