package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
        panic("Could not connect to database" + err.Error());
    }

	err = DB.Ping()
	if err != nil {
        panic("Could not ping the database" + err.Error());
    }

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTables()
	if err != nil {
        panic("Could not create tables" + err.Error());
    }
}

func createTables() error {
	createUsersTable := `
	    CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
        panic("Could not create the users table")
    }

	createEventsTable := `
	    CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
            user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
        panic("Could not create the events table")
    }

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create the registrations table")
	}
	

	return err
}