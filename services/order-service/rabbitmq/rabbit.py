from pydantic import BaseModel
import pika
class Rabbit(BaseModel):
    host:str
    user:str
    password:str
    port: int

    # def __init__(self, host, user, password):
    #     self.host=host
    #     self.password=password
    #     self.user=user
    def conection(self):
        connection = pika.BlockingConnection(pika.URLParameters(f'amqp://{self.user}:{self.password}@{self.host}:{self.port}/'))
        
        return connection