package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"

	maxEventsToInsert = 100000
	numberOfWorkers   = 50
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

func generateRandomEvents(db *sql.DB, countChan chan int) {
	for idx := range countChan {
		eventType := eventTypes[rand.Intn(len(eventTypes))]
		eventDescription := fmt.Sprintf("This is a random event %d", idx)
		eventApplication := eventApplications[rand.Intn(len(eventApplications))]

		_, err := insertLogEvent(db, eventType, eventDescription, eventApplication, generateRandomDate())
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}

	}
}

func generateRandomDate() time.Time {
	year := time.Now().Year()
	firstDayOfTheYear := time.Date(year, 1, 10, 0, 0, 0, 0, time.UTC)
	return firstDayOfTheYear.AddDate(0, rand.Intn(12), 0)
}
