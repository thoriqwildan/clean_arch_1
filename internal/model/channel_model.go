package model

type CreateChannelRequest struct {
	Name				string `json:"name" validate:"required"`
	Code 				string `json:"code,omitempty" validate:"required"`
	PaymentMethodId uint   `json:"paymentMethodId" validate:"required"`
	IconUrl			string `json:"iconUrl,omitempty"`
	OrderNum		int    `json:"orderNum,omitempty" validate:"omitempty,numeric"`
	LibName			string `json:"libName,omitempty"`
	Mdr 				int `json:"mdr,omitempty" validate:"omitempty,numeric"`
	FixedFee		float64 `json:"fixedFee,omitempty" validate:"omitempty,numeric"`
	UserAction	string `json:"userAction" validate:"required"`
}

type ChannelResponse struct {
	Id                uint    `json:"id"`
	Name              string  `json:"name"`
	PaymentMethod 		PaymentMethod  `json:"paymentMethod"`
	Code              string  `json:"code,omitempty"`
	IconUrl           string  `json:"iconUrl,omitempty"`
	OrderNum          int     `json:"orderNum,omitempty"`
	LibName           string  `json:"libName,omitempty"`
	Mdr               string  `json:"mdr,omitempty"`
	FixedFee          float64 `json:"fixedFee,omitempty"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}

type PaymentMethod struct {
	Id  uint   `json:"id"`
	Code string `json:"code"`
}