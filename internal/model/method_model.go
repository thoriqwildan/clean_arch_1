package model

type CreateMethodRequest struct {
	Name				string `json:"name" validate:"required"`
	Desc				string `json:"desc,omitempty"`
	OrderNum		int    `json:"orderNum,omitempty" validate:"omitempty,numeric"`
	UserAction	string `json:"userAction" validate:"required"`
	Code				string `json:"code,omitempty" validate:"required"`
}

type UpdateMethodRequest struct {
	ID					string   `json:"-" validate:"required,numeric"`
	Name				string `json:"name" validate:"required"`
	Desc				string `json:"desc,omitempty"`
	OrderNum		int    `json:"orderNum,omitempty" validate:"omitempty,numeric"`
	UserAction	string `json:"userAction" validate:"required"`
	Code				string `json:"code,omitempty" validate:"required"`
}

type MethodResponse struct {
	ID					uint   `json:"id"`
	Name				string `json:"name" validate:"required"`
	Desc				string `json:"desc,omitempty"`
	OrderNum		int    `json:"orderNum,omitempty" validate:"omitempty,numeric"`
	UserAction	string `json:"userAction" validate:"required"`
	Code				string `json:"code,omitempty" validate:"required"`
	CreatedAt		string `json:"createdAt"`
	UpdatedAt		string `json:"updatedAt"`
}

type FilterMethodRequest struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Page int `json:"page,omitempty" validate:"omitempty,numeric"`
	Limit int `json:"limit,omitempty" validate:"omitempty,numeric"`
}