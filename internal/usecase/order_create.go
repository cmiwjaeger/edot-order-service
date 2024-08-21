package usecase

import (
	"context"
	"edot-monorepo/services/order-service/internal/model"
	"edot-monorepo/services/order-service/internal/model/converter"

	"github.com/gofiber/fiber/v2"
)

type OrderCreateUseCase struct {
	*OrderBaseUseCase
}

func NewOrderCreateUseCase(orderBaseUseCase *OrderBaseUseCase) *OrderCreateUseCase {
	return &OrderCreateUseCase{
		orderBaseUseCase,
	}
}

func (c *OrderCreateUseCase) Exec(ctx context.Context, request *model.OrderCreateRequest) (*model.OrderResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.OrderRepository.Create(tx, nil); err != nil {
		c.Log.Warnf("Failed create order to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(nil), nil
}
