// mylocalproject github.com/tomhaerter/course-recommender

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tomhaerter/course-recommender/internal/database"
)

func main() {
	// Initialize your sqlc database client
	sqlClient := database.NewClient()
	defer sqlClient.Db.Close()

	applyDBMigrations(sqlClient)
	seedDBWithDummyData(sqlClient)
}

// Method declaration for applying migrations to the database
func applyDBMigrations(sqlClient *database.Db) {
	ctx := context.Background()

	// Directory where your migration files are located
	migrationsDir := "internal/db/migrations"

	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Could not list migration files: %v", err)
	}

	// Convert DirEntry slice to FileInfo slice to sort by Name
	var files []os.FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatalf("Could not get info for file: %v", err)
		}
		files = append(files, info)
	}

	// Sort files by name to ensure correct order
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") { // Ensure it's an SQL file
			log.Printf("Applying migration file: %s", file.Name())
			content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				log.Fatalf("Could not read migration file %s: %v", file.Name(), err)
			}

			sql := string(content)
			if _, err := sqlClient.Db.Exec(ctx, sql); err != nil {
				// log.Fatalf("Could not apply migration from file %s: %v", file.Name(), err)
				log.Printf("Could not apply migration from file %s: %v", file.Name(), err)
			}
		}
	}

	log.Println("All migrations applied successfully!")

}

// Method declaration for seeding the database with dummy data
func seedDBWithDummyData(sqlClient *database.Db) {
	ctx := context.Background()

	// SEED DB WITH DUMMY DATA
	// graph ql create new todo

	firstTodo, err := sqlClient.CreateTodo(ctx, "Setup seed")
	if err != nil {
		log.Fatalf("Could not create first todo: %v", err)

	}
	fmt.Printf("Created new todo: %+v\n", firstTodo)

	log.Printf("Finished running seed script")
}
