from pydantic import BaseModel
from typing import Optional
import pika
import json
import uuid
class Order(BaseModel):
    order_id: Optional[str]=None
    product_id: int
    quantity: int
    description: Optional[str]=None
    value: float

    def send_to_queue(self, order_data, connection):
        order_data["order_id"]=self._generate_id()
        print("iniciando envio da ordem de ID:  ", order_data["order_id"])
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
    
    def _generate_id(self):
        return  uuid.uuid4().hex