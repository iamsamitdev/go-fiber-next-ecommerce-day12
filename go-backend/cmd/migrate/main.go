package main

import (
	"flag"
	"log"
	"os"

	"github.com/iamsamitdev/fiber-ecommerce-api/internal/config"
)

func main() {
	var (
		up   = flag.Bool("up", false, "Run migrations")
		down = flag.Bool("down", false, "Rollback migrations (not implemented yet)")
	)
	flag.Parse()

	if !*up && !*down {
		log.Println("Usage:")
		log.Println("  go run cmd/migrate/main.go -up    # Run migrations")
		log.Println("  go run cmd/migrate/main.go -down  # Rollback migrations")
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup database connection
	db := config.SetupDatabase(cfg)

	if *up {
		log.Println("Running database migrations...")
		err := config.RunMigrationManual(cfg)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migration completed successfully!")

		// Seed database after migration
		log.Println("Seeding database...")
		if err := config.SeedDatabase(db, cfg); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}
		log.Println("🎉 All database operations completed successfully!")
	}

	if *down {
		log.Println("Rollback migrations not implemented yet")
		// TODO: Implement rollback functionality
	}
}
