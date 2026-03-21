# Nome do arquivo compose (opcional)
COMPOSE=docker compose

# Subir os containers
up:
	$(COMPOSE) up -d

# Derrubar os containers
down:
	$(COMPOSE) down

# Buildar as imagens
build:
	$(COMPOSE) build

# Ver logs
logs:
	$(COMPOSE) logs -f

# Reiniciar tudo
restart:
	$(COMPOSE) down && $(COMPOSE) up -d

# Executar comando dentro de um container
exec:
	$(COMPOSE) exec app sh

# Limpar tudo (containers, volumes, imagens)
clean:
	$(COMPOSE) down -v --rmi all --remove-orphans