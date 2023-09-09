package main

import (
	"encoding/json"
	"io"

	// "log"
	"fmt"

	"database/sql"
	"net/http"
	"os"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	apiUri = ":8080"
)

func main() {
	executeMigrations()

	go processQueue()

	http.HandleFunc("/v1/transfers", transfer)
	logger("main").Error(http.ListenAndServe(apiUri, nil))
}

func getDatabaseConnection() *sql.DB {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable")
	if err != nil {
		logger("database").Error(err)
	}
	err = conn.Ping()
	if err != nil {
		logger("database").Error(err)
	}
	return conn
}

func logger(context string) *log.Entry {
	log.SetFormatter(&log.TextFormatter{})
	return log.WithFields(log.Fields{"context": context})
}

func executeMigrations() {
	conn := getDatabaseConnection()
	if conn == nil {
		logger("migration").Fatal("Invalid connection created")
	}
	defer conn.Close()

	logger("migration").Info("Initialized connection")
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		logger("migration").Fatal(err)
	}
	m.Up()
}

type Transfer struct {
	Amount int `json:"amount" binding:"required"`
}

type Amount struct {
	Id     string `json:"id"`
	Amount int    `json:"amount"`
}

func transfer(w http.ResponseWriter, req *http.Request) {
	var transfer Amount
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&transfer)

	if err != nil {
		logger("transfer_json_parse").Error(err, req.Body)
		http.Error(w, "Invalid json provided", http.StatusUnprocessableEntity)
		return
	}

	transfer.Id = uuid.New().String()

	res, err := json.Marshal(transfer)
	if err != nil {
		logger("transfer").Error(err)
	}

	message := amqp.Publishing{
		Body:        res,
		ContentType: "application/json",
	}

	channelRabbitMq := getRabbitConnection()
	defer channelRabbitMq.Close()
	if err := channelRabbitMq.Publish(
		"",
		"pong",
		false,
		false,
		message,
	); err != nil {
		panic(err)
	}

	io.WriteString(w, "OK!\n")
}

func getRabbitConnection() *amqp.Channel {
	amqpUrl := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMq, err := amqp.Dial(amqpUrl)
	if err != nil {
		panic(err)
	}

	channelRabbitMq, err := connectRabbitMq.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channelRabbitMq.QueueDeclare(
		"pong", // queue name
		true,   // durable
		false,  // auto delete
		false,  // exclusive
		false,  // no wait
		nil,    // arguments
	)
	if err != nil {
		panic(err)
	}

	return channelRabbitMq
}

func processQueue() {
	connectRabbitMQ := getRabbitConnection()
	defer connectRabbitMQ.Close()

	messages, err := connectRabbitMQ.Consume(
		"pong", // queue name
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // arguments
	)
	if err != nil {
		logger("queue").Error(err)
	}
	logger("queue").Info("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			logger("queue").Info(fmt.Sprintf("Received message"))

			var data Amount
			if err := json.Unmarshal(message.Body, &data); err != nil {
				logger("queue").Error(err)
			}

			db := getDatabaseConnection()
			defer db.Close()

			sqlStatement := `
            INSERT INTO transfers (id, amount)
            VALUES ($1, $2)
            RETURNING id`
			id := ""
			err = db.QueryRow(sqlStatement, data.Id, data.Amount).Scan(&id)
			if err != nil {
				logger("queue").Error(err)
			}

			logger("queue").Info("New record ID is:", id)
		}
	}()

	<-forever
}
