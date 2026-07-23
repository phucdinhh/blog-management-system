package response

type Response[T any] struct {
	Data T `json:"data"`
}

type ListResponse[T any] struct {
	Data []T        `json:"data"`
	Meta Pagination `json:"meta"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}
