
from fastapi import FastAPI, HTTPException
from models import order
import uvicorn
from rabbitmq import rabbit
app = FastAPI()


broker=rabbit.Rabbit(user="guest", password="guest", host="localhost", port=5672)


@app.get("/")
async def index():
    return {"message":"welcome"}

@app.post("/checkout")
async def checkout(order:order.Order):
    connection=broker.conection()
    print(order)
    
    try:
        order.send_to_queue(order.dict(), connection)
        return {"status": "success", "message": "Pedido enviado para processamento"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    

if __name__ == "__main__":
    uvicorn.run("main:app", host="127.0.0.1", port=5000, reload=True)

