package rabbitmq

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbit struct {
	User     string
	Password string
	Host     string
	Port     int
	Conn     *amqp091.Connection
}

func NewRabbit(user string, port int, pass string, host string) *rabbit {

	return &rabbit{
		User:     user,
		Password: pass,
		Host:     host,
		Port:     port,
		Conn:     nil,
	}
}

func (r *rabbit) Conection() error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		r.User,
		r.Password,
		r.Host,
		r.Port,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("falha ao conectar no RabbitMQ em %s:%d: %w", r.Host, r.Port, err)
	}

	r.Conn = conn

	log.Printf("Conectado ao RabbitMQ em %s:%d", r.Host, r.Port)

	return nil
}
