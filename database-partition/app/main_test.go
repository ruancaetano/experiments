package main

import (
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	db := connectToDb()
	defer db.Close()

	referenceJanuaryDay := time.Date(time.Now().Year(), 1, 15, 0, 0, 0, 0, time.UTC)
	referenceFebruaryDay := referenceJanuaryDay.Add(time.Hour * 1 * 24 * 31)
	referenceMarchDay := referenceJanuaryDay.Add(time.Hour * 2 * 24 * 31)

	type args struct {
		eventType        string
		eventDescription string
		eventApplication string
		eventTime        time.Time
	}

	tests := []struct {
		tag               string
		args              args
		expectedPartition int
	}{
		{
			tag: "TestInsert partition 1",
			args: args{
				eventType:        "error",
				eventDescription: "This is a random event 3",
				eventApplication: "test3",
				eventTime:        referenceJanuaryDay,
			},
			expectedPartition: 1,
		},
		{
			tag: "TestInsert partition 2",
			args: args{
				eventType:        "error",
				eventDescription: "This is a random event 1",
				eventApplication: "test1",
				eventTime:        referenceFebruaryDay,
			},
			expectedPartition: 2,
		},
		{
			tag: "TestInsert partition 3",
			args: args{
				eventType:        "error",
				eventDescription: "This is a random event 2",
				eventApplication: "test2",
				eventTime:        referenceMarchDay,
			},
			expectedPartition: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			id, err := insertLogEvent(db, tt.args.eventType, tt.args.eventDescription, tt.args.eventApplication, tt.args.eventTime)
			if err != nil {
				t.Errorf("Error inserting log event: %v", err)
			}

			partitionExpectedTable := fmt.Sprintf("%s%d", partitionTablePrefix, tt.expectedPartition)
			selectResult, err := selectLogEvent(db, partitionExpectedTable, id)
			if err != nil {
				t.Errorf("Error selecting log event: %v", err)
			}

			if !selectResult.Next() {
				t.Errorf("No rows returned")
			}
		})
	}
}
