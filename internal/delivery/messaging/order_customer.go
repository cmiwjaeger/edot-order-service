package messaging

import (
	"edot-monorepo/services/order-service/internal/model"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type OrderConsumer struct {
	Log *logrus.Logger
}

func NewOrderConsumer(log *logrus.Logger) *OrderConsumer {
	return &OrderConsumer{
		Log: log,
	}
}

func (c OrderConsumer) Consume(message *kafka.Message) error {
	ContactEvent := new(model.Order)
	if err := json.Unmarshal(message.Value, ContactEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Contact event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic contacts with event: %v from partition %d", ContactEvent, message.TopicPartition.Partition)
	return nil
}
