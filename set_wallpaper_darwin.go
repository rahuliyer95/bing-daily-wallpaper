package main

import (
	"database/sql"
	"log"
	"os"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

// GetWallpaperPath returns the path to store the wallpaper
func GetWallpaperPath() string {
	return os.Getenv("HOME") + "/Pictures/wallpaper.jpg"
}

// SetWallpaper sets the wallpaper from path for macOS
func SetWallpaper(path string) {
	dbPath := os.Getenv("HOME") + "/Library/Application Support/Dock/desktoppicture.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Unable to open db at path %s.\nError is: %v\n", dbPath, err)
	}

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Unable to begin transaction.\nError is: %v\n", err)
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(`
	UPDATE data 
	SET value=? 
	WHERE ROWID IN (
		SELECT ROWID
		FROM data
		ORDER BY rowid ASC
	)
	`)
	if err != nil {
		log.Fatalf("Unable to prepare statement.\nError is: %v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(path)
	if err != nil {
		log.Fatalf("Unable to execute statement.\nError is: %v\n", err)
	}

	err = exec.Command("killall", "Dock").Run()
	if err != nil {
		log.Fatalf("Unable to kill dock.\nError is: %v\n", err)
	}

	log.Println("Wallpaper set successfully")
}
