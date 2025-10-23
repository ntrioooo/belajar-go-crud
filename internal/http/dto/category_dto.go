package dto

import "belajar-go/internal/core/domain"

type CategoryDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCategoryDTO(d *domain.Category) *CategoryDTO {
	if d == nil {
		return nil
	}
	return &CategoryDTO{
		ID:        d.ID,
		Name:      d.Name,
		CreatedAt: d.CreatedAt.Format(timeLayout),
		UpdatedAt: d.UpdatedAt.Format(timeLayout),
	}
}

func NewCategoryDTOs(ds []domain.Category) []CategoryDTO {
	out := make([]CategoryDTO, 0, len(ds))
	for i := range ds {
		out = append(out, *NewCategoryDTO(&ds[i]))
	}
	return out
}
