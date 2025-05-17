package commons

import "math"

const (
	DEFAULT_PAGE = 1
	DEFAULT_SIZE = 10
)

type Pagination struct {
	Page      int `json:"page" `
	Size      int `json:"size"`
	TotalPage int `json:"total"`
	TotalData int `json:"total_data"`
}

func NewPagination(page, size, totalData int) Pagination {
	return Pagination{
		Page:      page,
		Size:      size,
		TotalPage: int(math.Ceil(float64(totalData) / float64(size))),
		TotalData: totalData,
	}
}
