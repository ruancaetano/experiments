package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	mainTable            = "log_events"
	partitionTablePrefix = "log_events_month_"
	insertQuery          = `INSERT INTO log_events (event_type, event_description, event_application, event_time) VALUES ($1, $2, $3, $4) RETURNING event_id`
	selectQuery          = `SELECT * FROM %s WHERE event_id = $1`
)

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

func insertLogEvent(db *sql.DB, eventType string, eventDescription string, eventApplication string, eventTime time.Time) (int64, error) {
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(eventType, eventDescription, eventApplication, eventTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func selectLogEvent(db *sql.DB, table string, id int64) (*sql.Rows, error) {
	stmt, err := db.Prepare(fmt.Sprintf(selectQuery, table))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query(id)
}
