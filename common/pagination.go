package common

// Pagination is a struct for pagination
type Pagination struct {
	Page         int    `query:"page" json:"page" validate:"required"`
	PageSize     int    `query:"page_size" json:"page_size" validate:"required"`
	TotalRecords int    `json:"total_records"`
	TotalPages   int    `json:"total_pages"`
	MoreRecords  bool   `json:"more_records"`
	FirstPage    string `json:"first_page"`
	LastPage     string `json:"last_page"`
	NextPage     string `json:"next_page"`
	PrevPage     string `json:"prev_page"`
}
