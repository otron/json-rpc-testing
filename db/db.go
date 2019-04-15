package main

import (
	"log"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func main() {
	db := Database{db: Connect()}
	time.Sleep(5 * time.Second)
	db.Close()
}

func (db *Database) Close() {
	log.Println("Closing database connection.")
	db.db.Close()
}

func Connect () *sql.DB {
	addr := "root@tcp(172.26.0.2)/mbn"
	db, err := sql.Open("mysql", addr)
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err)
	}

	go func () {
		for {
			if err = db.Ping(); err != nil {
				log.Println("Could not ping database.", err)
				db.Close()
				db, err = sql.Open("mysql", addr)
			}
			time.Sleep(60 * time.Second)
		}
	}()
	return db
}
