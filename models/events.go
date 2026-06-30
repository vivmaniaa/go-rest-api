package models

import (
	"database/sql"
	"fmt"
	"time"

	"practice.com/rest-api/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	query := `INSERT INTO events( name, description, location, datetime, user_id)
	VALUES(?,?,?,?,?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Unable to prepare query for the saving!")
		return
	}
	defer db.DB.Close()
	var result sql.Result
	result, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		fmt.Println("Unable to execute the insert into events table")
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Unable to retrieve last insert id")
		return
	}
	e.ID = int(lastID)
	events = append(events, e)
}

func GetAllEvents() []Event {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Unable to get events from the 'events' table", err)
		return nil
	}
	defer rows.Close()
	defer db.DB.Close()
	var events []Event
	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		events = append(events, event)
	}

	return events
}
