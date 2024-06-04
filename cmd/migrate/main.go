package main

import (
	"garasystem/internal/adapters/postgrestorage"
	"garasystem/pkg/config"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	db, err := postgrestorage.NewConnection(postgrestorage.ParseFromConfig(cfg))

	if err != nil {
		log.Fatalf("cannot connecting to db: %v\n", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("cannot execute migration: %v\n", err)
	}

	total, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("cannot execute migration: %v\n", err)
	}

	log.Printf("applied %d migrations\n", total)
}
