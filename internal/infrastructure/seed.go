package infrastructure

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func RunSeed(db *sql.DB) {
	var count int
	// cek apakah admin sudah ada
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", "admin@gmail.com").Scan(&count)
	if err != nil {
		log.Println("Seed check failed:", err)
		return
	}

	if count > 0 {
		log.Println("Seed skipped (admin already exists)")
		return
	}

	// hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte("Semangatmuda123"), bcrypt.DefaultCost)

	_, err = db.Exec(`
		INSERT INTO users (name, email, password, role)
		VALUES (?, ?, ?, ?)
	`, "Superadmin", "admin@gmail.com", string(hashed), "superadmin")
	if err != nil {
		log.Println("Seed failed:", err)
		return
	}

	log.Println("Seed executed successfully")
}
