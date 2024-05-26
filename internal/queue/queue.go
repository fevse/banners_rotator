package queue

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Address    string
	Queue      string
	Exchange   string
	Kind       string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func New(address, queue, exchange, kind string) *Rabbit {
	return &Rabbit{Address: address, Queue: queue, Exchange: exchange, Kind: kind}
}

func (r *Rabbit) Connect() error {
	conn, err := amqp.Dial(r.Address)
	if err != nil {
		return fmt.Errorf("amqp dial error: %w", err)
	}
	r.Connection = conn
	ch, err := r.Connection.Channel()
	if err != nil {
		return fmt.Errorf("connection channer error: %w", err)
	}
	_, err = ch.QueueDeclare(
		r.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("queue declaration error: %w", err)
	}
	err = ch.ExchangeDeclare(
		r.Exchange,
		r.Kind,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("exchange declaration error: %w", err)
	}
	r.Channel = ch
	return nil
}

func (r *Rabbit) Close() error {
	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			return fmt.Errorf("channel close error: %w", err)
		}
	}
	if r.Connection != nil {
		if err := r.Connection.Close(); err != nil {
			return fmt.Errorf("connection close error: %w", err)
		}
	}
	return nil
}

func (r *Rabbit) Publish(msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Channel.PublishWithContext(
		ctx,
		r.Exchange,
		r.Queue,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(msg),
			DeliveryMode:    amqp.Transient,
		})
	if err != nil {
		return fmt.Errorf("publishing error: %w", err)
	}
	return nil
}
