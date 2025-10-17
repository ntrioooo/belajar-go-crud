package gormrepo

import "belajar-go/internal/core/domain"

func toDomainUser(m *User) *domain.User {
	if m == nil {
		return nil
	}
	return &domain.User{ID: m.ID, Email: m.Email, Password: m.Password, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}
}

func toModelUser(d *domain.User) *User {
	if d == nil {
		return nil
	}
	return &User{ID: d.ID, Email: d.Email, Password: d.Password, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt}
}

func toDomainPost(m *Post) *domain.Post {
	if m == nil {
		return nil
	}
	return &domain.Post{ID: m.ID, Title: m.Title, Body: m.Body, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}
}

func toModelPost(d *domain.Post) *Post {
	if d == nil {
		return nil
	}
	return &Post{ID: d.ID, Title: d.Title, Body: d.Body, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt}
}
