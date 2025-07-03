package converter

import (
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
)

func ChannelToResponse(channel *entity.PaymentChannel, method *entity.PaymentMethod) *model.ChannelResponse {
	return &model.ChannelResponse{
		Id: channel.ID,
		Name: channel.Name,
		PaymentMethod: model.PaymentMethod{
			Id: method.ID,
			Code: method.Code.String,
		},
		Code: channel.Code,
		IconUrl: channel.Code,
		OrderNum: int(channel.OrderNum.Int64),
		LibName: channel.LibName.String,
		Mdr: channel.MDR,
		FixedFee: channel.FixedFee,
		CreatedAt: channel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: channel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}