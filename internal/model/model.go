package model

type WebResponse[T any] struct {
	Success bool `json:"success"`
	Data T `json:"data,omitempty"`
	Meta *PaginationPage `json:"meta,omitempty"`
	Error any `json:"error,omitempty"`
}

type PaginationPage struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int64 `json:"total"`
}