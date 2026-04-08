package consumer

import (
	"context"
	"encoding/json"
	"inventory/database"
	"inventory/models"
	"inventory/repository"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	Conn   *amqp091.Channel
	DBConn *database.DBInventory
}

func NewConsumer(c *amqp091.Channel, db *database.DBInventory) *consumer {

	return &consumer{
		Conn:   c,
		DBConn: db,
	}
}

func (c *consumer) Read() error {
	repo := repository.NewProductRepo(c.DBConn)
	//var product models.Product
	var order models.Order
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
			err := json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Printf("Erro ao decodificar mensagem: %s", err)
				continue
			}
			if _, err := repo.GetProduct(context.Background(), order.ProductID); err != nil {
				log.Printf("The product does't exists")
				continue
			}

			err = repo.ReduceStock(context.Background(), order.ProductID)

			if err != nil {
				log.Printf("Failure on do the order: %v", err)
			} else {

				log.Println("Product quantitty decrease from product ID ")
			}

		}
	}()

	log.Printf(" [*] Aguardando mensagens de pedidos. Para sair pressione CTRL+C")
	<-forever

	return nil
}
