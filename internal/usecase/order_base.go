package usecase

import (
	repository "edot-monorepo/services/order-service/internal/repository/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderBaseUseCase struct {
	DB              *gorm.DB
	Log             *logrus.Logger
	OrderRepository *repository.OrderRepository
	Validate        *validator.Validate
}

func NewOrderUseCase(db *gorm.DB, log *logrus.Logger, orderRepo *repository.OrderRepository, validate *validator.Validate) *OrderBaseUseCase {
	return &OrderBaseUseCase{
		DB:              db,
		Log:             log,
		OrderRepository: orderRepo,
		Validate:        validate,
	}
}
