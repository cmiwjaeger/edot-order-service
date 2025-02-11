package repository

import (
	"edot-monorepo/services/order-service/internal/entity"

	"github.com/sirupsen/logrus"
)

type OrderRepository struct {
	Repository[entity.Order]
	Log *logrus.Logger
}

func NewOrderRepository(log *logrus.Logger) *OrderRepository {
	return &OrderRepository{
		Log: log,
	}
}
