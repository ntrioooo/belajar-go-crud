package dto

import "time"

// Satu tempat untuk format waktu agar konsisten di semua response.
// RFC3339 (UTC) atau lokal — di contoh pakai RFC3339.
const timeLayout = time.RFC3339

// (Opsional) Envelope standar jika butuh meta/pagination:
type PageMeta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}
