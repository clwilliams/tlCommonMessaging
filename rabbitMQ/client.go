package rabbitMQ

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

// MessageClient - wrapper for rabbitMQ
type MessageClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// Connect - connects to the message broker
func (messageClient *MessageClient) Connect(
	rabbitMqHost, rabbitMqPort, rabbitMqUser, rabbitMqPassword *string) error {
	var err error
	messageClient.Connection, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		*rabbitMqUser, *rabbitMqPassword, *rabbitMqHost, *rabbitMqPort))
	if err != nil {
		return fmt.Errorf("Problem connecting to RabbitMQ %v", err)
	}
	return nil
}

// ConfigureChannelAndExchange - creates a channel and configures to given exchange
func (messageClient *MessageClient) ConfigureChannelAndExchange(rabbitMqExchange *string) error {
	// setup channel
	var err error
	messageClient.Channel, err = messageClient.Connection.Channel()
	if err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ Channel %v", err)
	}
	// setup exchange
	if err = messageClient.Channel.ExchangeDeclare(*rabbitMqExchange, "topic", true, false, false, false, nil); err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ Topic %v", err)
	}
	return nil
}

// SendMessage - posts a message to RabbitMQ
func (messageClient *MessageClient) SendMessage(exchange, routingKey, body string) error {
	ch, err := messageClient.Connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err := ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
			MessageId:   time.Now().String(),
		}); err != nil {
		return fmt.Errorf("Problem sending message to RabbitMQ %v", err)
	}
	return nil
}
