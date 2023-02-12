package requests

import (
	"ganhaum.henrybarreto.dev/internal/api/http/queries"
)

type ProductCreate struct {
	Thumbnail   string `json:"thumbnail" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Kind        string `json:"type" binding:"required"`
	Value       int    `json:"value" binding:"required"`
}
type ProductGet struct {
	ID int `uri:"id" binding:"required"`
}

type ProductsGet struct {
	queries.Pagination
}

type ProductUpdate struct {
	ID          int    `uri:"id" binding:"required"`
	Thumbnail   string `json:"thumbnail" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Kind        string `json:"type" binding:"required"`
	Value       int    `json:"value" binding:"required"`
}

type ProductDelete struct {
	ID int `uri:"id" binding:"required"`
}
