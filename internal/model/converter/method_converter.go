package converter

import (
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
)

func MethodToResponse(method *entity.PaymentMethod) *model.MethodResponse {
	return &model.MethodResponse{
		ID: method.ID,
		Name: method.Name,
		Desc: method.Desc.String,
		OrderNum: method.OrderNum,
		UserAction: method.UserAction,
		Code: method.Code.String,
		CreatedAt: method.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: method.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}