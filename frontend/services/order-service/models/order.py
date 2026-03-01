from pydantic import BaseModel
import pika
import json
class Order(BaseModel):
    product_id: int
    quantity: int


def send_to_queue(order_data):
    # 'rabbitmq' é o nome do serviço definido no docker-compose
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='rabbitmq'))
    channel = connection.channel()

    # Garante que a fila existe
    channel.queue_declare(queue='order_queue', durable=True)

    message = json.dumps(order_data)
    channel.basic_publish(
        exchange='',
        routing_key='order_queue',
        body=message,
        properties=pika.BasicProperties(
            delivery_mode=2,  # Torna a mensagem persistente
        ))
    connection.close()