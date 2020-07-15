package blogs

import "database/sql"

type Blog struct {
	Id        int64          `json:"id_bundle"`
	Title     string         `json:"blog_title"`
	Body      string         `json:"blog_body"`
	IsActive  int            `json:"is_active"`
	CreatedAt sql.NullString `json:"created_at"`
	CreatedBy int64          `json:"created_by"`
	UpdatedAt string         `json:"updated_at"`
	UpdatedBy int64          `json"updated_by"`
	DeletedAt string         `json:"deleted_at"`
	DeletedBy int64          `json:"deleted_by"`
}

type (
	Blogs []Blog
)
