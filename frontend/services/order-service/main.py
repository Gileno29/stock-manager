
from fastapi import FastAPI, HTTPException
from models import order
app = FastAPI()




@app.post("/checkout")
async def checkout(order:order.Order):
    try:
        order.send_to_queue(order.dict())
        return {"status": "success", "message": "Pedido enviado para processamento"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))