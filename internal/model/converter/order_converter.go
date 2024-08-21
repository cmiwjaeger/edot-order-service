package converter

import (
	"edot-monorepo/services/order-service/internal/entity"
	"edot-monorepo/services/order-service/internal/model"
)

func OrderToResponse(user *entity.Order) *model.OrderResponse {
	return &model.OrderResponse{
		ID: user.ID,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func OrderToTokenResponse(user *entity.Order) *model.OrderResponse {
	return &model.OrderResponse{}
}
