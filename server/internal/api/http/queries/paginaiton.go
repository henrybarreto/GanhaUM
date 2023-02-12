package queries

type Pagination struct {
	Page  int `form:"page" binding:"required,min=1"`
	Limit int `form:"limit" binding:"required,min=1,max=100"`
}
