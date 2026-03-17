package infrastructure

import (
	"database/sql"
	"log"
	"time"
)

func WaitForDB(db *sql.DB) {
	for i := 0; i < 10; i++ {
		err := db.Ping()
		if err == nil {
			log.Println("Database connected")
			return
		}
		log.Println("Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Database not ready")
}