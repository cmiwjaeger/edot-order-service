package config

import (
	"edot-monorepo/services/order-service/internal/delivery/http/controller"
	"edot-monorepo/services/order-service/internal/delivery/http/route"
	repository "edot-monorepo/services/order-service/internal/repository/gorm"
	"edot-monorepo/services/order-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	orderRepository := repository.NewOrderRepository(config.Log)
	orderBaseUseCase := usecase.NewOrderUseCase(config.DB, config.Log, orderRepository, config.Validate)
	orderCreateUseCase := usecase.NewOrderCreateUseCase(orderBaseUseCase)

	orderController := controller.NewOrderController(orderCreateUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:             config.App,
		OrderController: orderController,
	}

	routeConfig.Setup()
}
