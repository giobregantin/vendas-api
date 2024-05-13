package produtos

import (
	"encoding/json"

	"github.com/hsxflowers/vendas-api/exceptions"
	"github.com/hsxflowers/vendas-api/pkg/broker"
	"github.com/hsxflowers/vendas-api/produtos/domain"
	"github.com/labstack/gommon/log"
)

type ServiceConsumer interface {
	ReadMessagePagamento(topico string) (*domain.ConfirmacaoPagamento, error)
	SendEvent(topico string, key string, requestId string, payload interface{}) error
	ReadMessageRastreio(topico string) (*domain.CodigoRastreio, error)
}

type Consumer struct {
	ProdutosBroker *broker.Broker
}

func NewConsumer(produtosBroker *broker.Broker) ServiceConsumer {
	return &Consumer{produtosBroker}
}

func (c *Consumer) ReadMessagePagamento(topico string) (*domain.ConfirmacaoPagamento, error) {
	event, err := c.ProdutosBroker.Consumer.ReadEvent(topico)
	if err != nil {
		c.ProdutosBroker.Consumer.CommitMessage(event.Event)
		return nil, exceptions.New(exceptions.ErrReadingEvent, err)
	}

	message, err := ParseEventPagamento(event)
	if err != nil {
		c.ProdutosBroker.Consumer.CommitMessage(event.Event)
		return nil, err
	}

	c.ProdutosBroker.Consumer.CommitMessage(event.Event)
	return message, nil
}

func ParseEventPagamento(msg *broker.Event) (*domain.ConfirmacaoPagamento, error) {
	message := domain.ConfirmacaoPagamento{}

	err := json.Unmarshal(msg.Values, &message)
	if err != nil {
		log.Error("[Read Message] error parsing event", err)
		return nil, exceptions.New(exceptions.ErrUnprocessableJson, err)
	}
	return &message, nil

}

func (c *Consumer) ReadMessageRastreio(topico string) (*domain.CodigoRastreio, error) {
	event, err := c.ProdutosBroker.Consumer.ReadEvent(topico)
	if err != nil {
		c.ProdutosBroker.Consumer.CommitMessage(event.Event)
		return nil, exceptions.New(exceptions.ErrReadingEvent, err)
	}

	message, err := ParseEventRastreio(event)
	if err != nil {
		c.ProdutosBroker.Consumer.CommitMessage(event.Event)
		return nil, err
	}

	c.ProdutosBroker.Consumer.CommitMessage(event.Event)
	return message, nil
}

func ParseEventRastreio(msg *broker.Event) (*domain.CodigoRastreio, error) {
	message := domain.CodigoRastreio{}

	err := json.Unmarshal(msg.Values, &message)
	if err != nil {
		log.Error("[Read Message] error parsing event", err)
		return nil, exceptions.New(exceptions.ErrUnprocessableJson, err)
	}
	return &message, nil

}

func (c *Consumer) SendEvent(topico string, key string, requestId string, payload interface{}) error {
	err := c.ProdutosBroker.Producer.SendEvent(topico, key, requestId, payload)
	if err != nil {
		return exceptions.New(exceptions.ErrSendEvent, err)
	}

	return nil
}
