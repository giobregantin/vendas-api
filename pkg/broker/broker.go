package broker

type Broker struct {
	Consumer Consumer
	Producer Producer
}

type Event struct {
	Headers map[string]string
	Values  []byte
	Event   interface{}
}

type Consumer interface {
	ReadEvent(topic string) (*Event, error)
	CommitMessage(event interface{}) error
}

type Producer interface {
	SendEvent(topico string, key string, requestId string, body interface{}) error
}

func NewBroker(consumer Consumer, producer Producer) *Broker {
	return &Broker{
		Consumer: consumer,
		Producer: producer,
	}
}
