build-db:
		docker build -t pg .

run-db:
	docker run --name pg --rm -d -p 5432:5432 pg

stop-db:
	docker rm pg -f

run:
	go run main.go