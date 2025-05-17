package commons

const (
	DEFAULT_PAGE = 1
	DEFAULT_SIZE = 10
)

type Pagination struct {
	Page      int    `json:"page" `
	Size      int    `json:"size"`
	TotalPage int    `json:"total"`
	NextUrl   string `json:"next_url"`
	PrevUrl   string `json:"prev_url"`
}

func NewPagination(page, size, total int) Pagination {
	if page < 1 {
		page = DEFAULT_PAGE
	}
	if size < 1 {
		size = DEFAULT_SIZE
	}
	return Pagination{
		Page:      page,
		Size:      size,
		TotalPage: total,
	}
}
