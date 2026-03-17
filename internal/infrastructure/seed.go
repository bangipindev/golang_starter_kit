package infrastructure

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func RunSeed(db *sql.DB) {
	var role string
	var count int

	// ===============================
	// 1. Seed Role
	// ===============================
	err := db.QueryRow("SELECT id FROM roles WHERE name = ?", "superadmin").Scan(&role)

	if err == sql.ErrNoRows {
		_, err := db.Exec("INSERT INTO roles (name,guard_name) VALUES (?,?)", "superadmin", "web")
		if err != nil {
			log.Println("Seed role failed:", err)
			return
		}

		log.Println("Role superadmin created")
	} else if err != nil {
		log.Println("Role check failed:", err)
		return
	} else {
		log.Println("Role already exists")
	}

	// ===============================
	// 2. Seed User
	// ===============================
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", "admin@gmail.com").Scan(&count)
	if err != nil {
		log.Println("User check failed:", err)
		return
	}

	if count > 0 {
		log.Println("User already exists, skip seeding")
		return
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte("Semangatmuda123"), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Hash password failed:", err)
		return
	}

	_, err = db.Exec(`
		INSERT INTO users (name, email, password, role)
		VALUES (?, ?, ?, ?)
	`, "Superadmin", "admin@gmail.com", string(hashed), "superadmin")

	if err != nil {
		log.Println("Seed user failed:", err)
		return
	}

	log.Println("Seeder executed successfully")
}
