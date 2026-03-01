package main

import (
	"fmt"
	"inventory/consumer"
	"inventory/rabbitmq"
	"log"
)

func main() {
	rabbit := rabbitmq.NewRabbit("guest", 5672, "guest", "localhost")

	err := rabbit.Conection()
	if err != nil {
		fmt.Println("we have an error: ", err)
	}

	ch, err := rabbit.Conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir canal: %v", err)
	}
	defer ch.Close()
	defer rabbit.Conn.Close()

	consumer := consumer.NewConsumer(ch)

	consumer.Read()

	// Declara a fila (deve ser igual à do Python)

}
