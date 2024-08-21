package controller

import (
	"edot-monorepo/services/order-service/internal/model"
	"edot-monorepo/services/order-service/internal/usecase"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OrderController struct {
	orderCreateUseCase *usecase.OrderCreateUseCase
	Log                *logrus.Logger
	Validate           *validator.Validate
}

func NewOrderController(orderCreateUseCase *usecase.OrderCreateUseCase, log *logrus.Logger, validate *validator.Validate) *OrderController {
	return &OrderController{
		orderCreateUseCase: orderCreateUseCase,
		Log:                log,
		Validate:           validate,
	}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {

	return ctx.JSON(model.WebResponse[*model.OrderResponse]{Data: nil})
}
