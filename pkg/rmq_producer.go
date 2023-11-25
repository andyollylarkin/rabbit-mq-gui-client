package pkg

import "github.com/streadway/amqp"

func CreateProducer(uri string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}
