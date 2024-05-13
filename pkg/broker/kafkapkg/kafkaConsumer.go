package kafkapkg

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hsxflowers/vendas-api/pkg/broker"
)

type KafkaConsumerConfig struct {
	Broker     string
	GroupId    string
	AutoOffset string
	Topic      string
}

func NewKafkaConsumer(config *KafkaConsumerConfig) (*KafkaConsumerImp, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Broker,
		"group.id":          config.GroupId,
		"auto.offset.reset": config.AutoOffset,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaConsumerImp{
		Consumer: consumer,
		Topic:    config.Topic,
	}, nil
}

type KafkaConsumerImp struct {
	Consumer *kafka.Consumer
	Topic    string
}

func (c *KafkaConsumerImp) ReadEvent(topic string) (*broker.Event, error) {
	err := c.Consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	for {
		msg, err := c.Consumer.ReadMessage(-1)
		headers := make(map[string]string)
		for _, value := range msg.Headers {
			headers[value.Key] = string(value.Value)
		}
		brokerEvent := broker.Event{Headers: headers, Values: msg.Value, Event: msg}
		if err != nil {
			return nil, err
		} else {
			return &brokerEvent, nil
		}
	}
}

func (c *KafkaConsumerImp) CommitMessage(event interface{}) error {
	msg, ok := event.(*kafka.Message)
	if !ok {
		return nil
	}

	_, err := c.Consumer.CommitMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
