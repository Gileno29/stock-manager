package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"inventory/database"
	"inventory/models"
	"inventory/repository"
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
			db := database.NewDB("postgres", "user", "password", 5434, "localhost", "inventory_service_db")

			err = db.Conection()

			if err != nil {
				fmt.Println("erro ao conectar no banco de dados", err)
			}
			db.Populate()
			repo := repository.NewRepo(db)

			_, err = repo.InsertStock(context.Background(), order.ProductID, order.Quantity)

			if err != nil {
				log.Printf("Falha ao inserir no estoque: %v", err)
			} else {
				log.Printf("Estoque atualizado para o Produto %d", order.ProductID)
			}

		}
	}()

	log.Printf(" [*] Aguardando mensagens de pedidos. Para sair pressione CTRL+C")
	<-forever

	return nil
}
