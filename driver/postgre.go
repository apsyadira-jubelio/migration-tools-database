package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"sync"

	migrate "github.com/rubenv/sql-migrate"
)

func PostgreDbClient(host, port, user, password, database string) *sql.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatalf(fmt.Sprintf("error, not connected to database, %s", err.Error()))
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		log.Fatalf(fmt.Sprintf("error, not sent ping to database, %s", err.Error()))
	}

	return db
}

func CreatePostgreSchema(db *sql.DB, schemaName string, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Run migrations & create schema")

	schemaSql := []string{
		fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s AUTHORIZATION %s;`, schemaName, os.Getenv("DB_TENANT_USER")),
		fmt.Sprintf("SET search_path TO %s, pg_catalog;", schemaName),
	}

	for _, cmd := range schemaSql {
		stmt, err := db.Prepare(cmd)

		if err != nil {
			log.Fatalf("Error preparing statement: %v", err)
			return
		}

		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			log.Fatalf("Error executing statement: %v", err)
			return
		}
	}

	log.Printf("Schema %s created successfully", schemaName)
	MigrationDbTenant(db)
}

func DBSystemMigrate(db *sql.DB) {
	var migrations = &migrate.FileMigrationSource{
		Dir: "migrations/system",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error excute migration sql for system %s", err))
	} else {
		log.Println("Applied " + strconv.Itoa(n) + " migrations!")
	}
}

// Refactoring this code if database staging already setup
func MigrationDbTenant(db *sql.DB) {
	log.Println("prepare to migrate")
	var migrations = &migrate.FileMigrationSource{
		Dir: path.Join("./migrations/tenants"),
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error excute migration sql for tenant, %s", err.Error()))
	} else {
		log.Println("Applied " + strconv.Itoa(n) + " migrations!")
	}
}
