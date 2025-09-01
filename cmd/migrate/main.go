package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

	"lms-backend/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}

	// Connect database via GORM
	config.ConnectDB()

	// Pastikan tabel migrations ada
	err := config.DB.Exec(`CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE,
		batch INT NOT NULL,
		migrated_at TIMESTAMP DEFAULT NOW()
	);`).Error
	if err != nil {
		log.Fatal("Failed to create migrations table:", err)
	}

	if err := runMigrations(); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations completed successfully")
}

func runMigrations() error {
	// Ambil semua file .sql di folder migrations
	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		return err
	}
	sort.Strings(files)

	// Ambil semua migration yang sudah dijalankan
	executed := map[string]bool{}
	rows, err := config.DB.Raw(`SELECT name FROM migrations`).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name)
		executed[name] = true
	}

	batch := 1
	for _, file := range files {
		name := filepath.Base(file)
		if executed[name] {
			continue // skip migration yang sudah dijalankan
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %v", name, err)
		}

		if err := config.DB.Exec(string(content)).Error; err != nil {
			return fmt.Errorf("failed migration %s: %v", name, err)
		}

		if err := config.DB.Exec(`INSERT INTO migrations(name, batch) VALUES(?, ?)`, name, batch).Error; err != nil {
			return fmt.Errorf("failed to log migration %s: %v", name, err)
		}

		log.Printf("Migrated: %s\n", name)
	}

	return nil
}
