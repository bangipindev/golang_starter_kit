package infrastructure

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RunSeed(db *sql.DB) {
	var roleID int64
	var count int

	// ===============================
	// 1. Seed Role
	// ===============================
	err := db.QueryRow("SELECT id FROM roles WHERE name = ?", "superadmin").Scan(&roleID)

	if err == sql.ErrNoRows {
		res, err := db.Exec("INSERT INTO roles (name,guard_name,created_at,updated_at) VALUES (?,?,?,?)", "superadmin", "web", time.Now(), time.Now())
		if err != nil {
			log.Println("Seed role failed:", err)
			return
		}
		roleID, _ = res.LastInsertId()
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

	res, err := db.Exec(`
		INSERT INTO users (name, email, password, public_id, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, "Superadmin", "admin@gmail.com", string(hashed), uuid.New().String(), 1, time.Now(), time.Now())

	if err != nil {
		log.Println("Seed user failed:", err)
		return
	}

	userID, err := res.LastInsertId()
	if err == nil {
		_, err = db.Exec("INSERT INTO model_has_roles (role_id, model_type, model_id) VALUES (?, 'User', ?)", roleID, userID)
		if err != nil {
			log.Println("Failed to assign superadmin role to seeded user:", err)
		}
	} else {
		log.Println("Failed to get last insert ID for user:", err)
	}

	log.Println("Seeder executed successfully")
}
