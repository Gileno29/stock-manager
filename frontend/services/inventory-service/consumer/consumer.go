package consumer

import (
	"encoding/json"
	"inventory/models"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	Conn *amqp091.Channel
}

func NewConsumer(c *amqp091.Channel) *consumer {

	return &consumer{
		Conn: c,
	}
}

func (c *consumer) Read() error {
	q, err := c.Conn.QueueDeclare(
		"order_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	if err != nil {
		return err
	}

	msgs, err := c.Conn.Consume(
		q.Name, "", true, false, false, false, nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var order models.Order
			err := json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Printf("Erro ao decodificar mensagem: %s", err)
				continue
			}

			log.Printf("📦 Processando Estoque: Removendo %d unidades do Produto ID %d", order.Quantity, order.ProductID)

		}
	}()

	log.Printf(" [*] Aguardando mensagens de pedidos. Para sair pressione CTRL+C")
	<-forever

	return nil
}
