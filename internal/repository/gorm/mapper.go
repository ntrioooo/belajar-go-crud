package gormrepo

import "belajar-go/internal/core/domain"

func toDomainUser(m *User) *domain.User {
	if m == nil {
		return nil
	}
	return &domain.User{
		ID: m.ID, Email: m.Email, Username: m.Username, Password: m.Password,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt,
	}
}

func toModelUser(d *domain.User) *User {
	if d == nil {
		return nil
	}
	return &User{
		ID: d.ID, Email: d.Email, Username: d.Username, Password: d.Password,
		CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt,
	}
}

func toDomainPost(m *Post) *domain.Post {
	if m == nil {
		return nil
	}
	return &domain.Post{ID: m.ID, Title: m.Title, Body: m.Body, UserID: m.UserID, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}
}

func toModelPost(d *domain.Post) *Post {
	if d == nil {
		return nil
	}
	return &Post{ID: d.ID, Title: d.Title, Body: d.Body, UserID: d.UserID, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt}
}

func toModelCategory(d *domain.Category) *Category {
	if d == nil {
		return nil
	}
	return &Category{ID: d.ID, Name: d.Name, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt}
}

func toDomainCategory(m *Category) *domain.Category {
	if m == nil {
		return nil
	}
	return &domain.Category{ID: m.ID, Name: m.Name, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}
}
