package model

type WebResponse[T any] struct {
	Data T `json:"data"`
	Meta *PaginationPage `json:"meta,omitempty"`
}

type PaginationPage struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int64 `json:"total"`
}