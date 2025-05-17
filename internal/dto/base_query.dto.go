package dto

type BaseQuery struct {
	Limit  int    `json:"limit" form:"limit"`
	Page   int    `json:"page" form:"page"`
	Search string `json:"search" form:"search"`
}
