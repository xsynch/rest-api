package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB 

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3","./api.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer DB.Close()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}


func createTables(){
	createUsersTable := `
		create table if not exists users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email Text not null unique,
			password Text not null
		)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {		
		panic("Could not create users table")
	}

	createEventsTable := ` 
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		foreign key(user_id) references users(id)
	)
	`

	// fmt.Printf("Creating the DB AND tables: %v\n", createEventsTable)

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create events table")
		
	}
	createRegistration := `
	create table if not exists registrations (
	id integer primary key autoincrement
	event_id integer,
	user_id integer,
	foreign key (event_id) references events(id),
	foreign key (user_id) references users(id)
	)`
	
	_, err = DB.Exec(createRegistration)
	if err != nil {
		fmt.Println(err)
		panic("Could not create Registration table")
		
	}

}