package kafkapkg

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type producerImp struct {
	producer *kafka.Producer
}

func NewKafkaProducer(broker string) (*producerImp, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})

	if err != nil {
		return nil, err
	}

	producerImpl := producerImp{
		producer: producer,
	}

	return &producerImpl, nil
}

func (p *producerImp) SendEvent(topic string, channelId string, requestId string, body interface{}) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	kafkaEvent := &kafka.Message{
		Value: payload,
		Key:   []byte(channelId),
		Headers: []kafka.Header{
			{
				Key:   "X-Request-Id",
				Value: []byte(requestId),
			},
		},
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: -1},
	}

	err = p.producer.Produce(kafkaEvent, nil)

	if err != nil {
		return nil
	}

	return err
}
