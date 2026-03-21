from pydantic import BaseModel
import pika
import json
class Order(BaseModel):
    product_id: int
    quantity: int

    def send_to_queue(self, order_data, connection):
        # Garante que a fila existe
        channel = connection.channel()
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
        