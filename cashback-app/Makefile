.PHONY: help

# Show this help.
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

live-api:
	cd app; AMQP_SERVER_URL="amqp://myuser:mypassword@localhost:5672" air -- -mode=api

live-worker:
	cd app; AMQP_SERVER_URL="amqp://myuser:mypassword@localhost:5672" air -- -mode=worker

live-migration:
	cd app; AMQP_SERVER_URL="amqp://myuser:mypassword@localhost:5672" air -- -mode=migration

# Execute migrations
run-migrations:
	cd app; go run main.go

run-migrate:
		migrate -source file://migrations/migrations -database postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable up

# Install migrate golang with brew
install-migrate:
		@brew install golang-migrate

# Create new migration for database
create-migration:
		cd migrations; migrate create -ext sql -dir migrations -seq $(NAME)

# Run loop with ping
#watch -n $(N) curl -X POST localhost:8080/v1/transfers -d '{"amount": 38293}'
run-loop:
	 curl -X POST localhost:8080/v1/transfers -d '{"amount": 38293}'

req:
	 curl -X POST localhost:8080/v1/transfers -d '{"amount": 38293}'

req-error:
		curl -X POST localhost:8080/v1/transfers -d '{"amout": 38293'

# Initialize databases
up:
		docker-compose up -d --build
	
up-db:
		docker-compose up db -d 
		docker-compose up rabbitmq -d 
	
down:
		docker-compose down -v
	
recreate: down up

ps:
		docker-compose ps

logs:
		docker-compose logs
	
logs-api:
		docker-compose logs api

logs-worker:
		docker-compose logs worker

