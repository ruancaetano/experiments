package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"

	maxEventsToInsert = 10000
	numberOfWorkers   = 100
	insertQuery       = `INSERT INTO log_events (event_type, event_description, event_application) VALUES ($1, $2, $3)`
)

var (
	eventTypes        = []string{"error", "info", "warning"}
	eventApplications = []string{"app1", "app2", "app3", "app4", "app5"}
)

func main() {
	db := connectToDb()
	defer db.Close()
	fmt.Println("Successfully connected to the database")

	countChan := make(chan int, numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		go generateRandomEvents(db, countChan)
	}

	totalLogs := 0
	startTime := time.Now()
	for {
		if totalLogs == maxEventsToInsert {
			break
		}
		countChan <- totalLogs
		totalLogs++
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Println("Inserted all log events")
	fmt.Println("Tempo de execução:", elapsedTime)
}

func connectToDb() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func generateRandomEvents(db *sql.DB, countChan chan int) {
	for idx := range countChan {
		eventType := eventTypes[rand.Intn(len(eventTypes))]
		eventDescription := fmt.Sprintf("This is a random event %d", idx)
		eventApplication := eventApplications[rand.Intn(len(eventApplications))]

		insertLogEvent(db, eventType, eventDescription, eventApplication)

	}
}

func insertLogEvent(db *sql.DB, eventType string, eventDescription string, eventApplication string) error {
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventType, eventDescription, eventApplication)
	if err != nil {
		return err
	}
	return nil
}
