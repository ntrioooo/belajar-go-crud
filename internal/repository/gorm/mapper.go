package gormrepo

import "belajar-go/internal/core/domain"

// ===== User =====
func toDomainUser(m *User) *domain.User {
	if m == nil {
		return nil
	}
	return &domain.User{
		ID: m.ID, Email: m.Email, Username: m.Username, Password: m.Password, Role: m.Role,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt,
	}
}
func toModelUser(d *domain.User) *User {
	if d == nil {
		return nil
	}
	return &User{
		ID: d.ID, Email: d.Email, Username: d.Username, Password: d.Password, Role: d.Role,
		CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt,
	}
}

// ===== Post =====
func toDomainPost(m *Post) *domain.Post {
	if m == nil {
		return nil
	}
	return &domain.Post{
		ID: m.ID, Title: m.Title, Body: m.Body, UserID: m.UserID, CategoryID: m.CategoryID,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt,
	}
}
func toModelPost(d *domain.Post) *Post {
	if d == nil {
		return nil
	}
	return &Post{
		ID: d.ID, Title: d.Title, Body: d.Body, UserID: d.UserID, CategoryID: d.CategoryID,
		CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt,
	}
}

// ===== Category =====
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

// ===== Comment =====
func toDomainComment(m *Comment) *domain.Comment {
	if m == nil {
		return nil
	}
	return &domain.Comment{
		ID: m.ID, PostID: m.PostID, UserID: m.UserID, ParentID: m.ParentID, Body: m.Body,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt,
	}
}
func toModelComment(d *domain.Comment) *Comment {
	if d == nil {
		return nil
	}
	return &Comment{
		ID: d.ID, PostID: d.PostID, UserID: d.UserID, ParentID: d.ParentID, Body: d.Body,
		CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt,
	}
}

// ===== Retweet =====
func toDomainRetweet(m *Retweet) *domain.Retweet {
	if m == nil {
		return nil
	}
	return &domain.Retweet{
		ID: m.ID, UserID: m.UserID, OriginalPostID: m.OriginalPostID, QuoteBody: m.QuoteBody,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt,
	}
}
func toModelRetweet(d *domain.Retweet) *Retweet {
	if d == nil {
		return nil
	}
	return &Retweet{
		ID: d.ID, UserID: d.UserID, OriginalPostID: d.OriginalPostID, QuoteBody: d.QuoteBody,
		CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt,
	}
}
