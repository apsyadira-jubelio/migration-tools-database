package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	migrate "github.com/rubenv/sql-migrate"
)

func PostgreDbClient(dataSource string) *sql.DB {
	db, err := sql.Open("postgres", dataSource)

	if err != nil {
		log.Fatalf(fmt.Sprintf("error, not connected to database, %s", err.Error()))
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(2 * time.Minute)

	return db
}

// ensureDbConnection attempts to ping the database to verify connection is alive, and retries if not.
func EnsureDbConnection(db *sql.DB, maxRetries int) bool {
	for i := 0; i < maxRetries; i++ {
		err := db.Ping()
		if err == nil {
			return true // Connection is alive
		}

		log.Printf("Failed to ping DB, retrying in 5 seconds... (%v)", err)
		time.Sleep(5 * time.Second)
	}

	return false
}

func CreatePostgreSchema(db *sql.DB, schemaName string) (err error) {
	log.Printf("Run migrations & create schema of %s", schemaName)

	schemaSql := []string{
		fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s AUTHORIZATION %s;`, schemaName, os.Getenv("DATABASE_TENANT_USERNAME")),
		fmt.Sprintf("SET search_path TO %s, pg_catalog;", schemaName),
	}

	for _, sqlCmd := range schemaSql {
		// Execute SQL command
		if _, err := db.Exec(sqlCmd); err != nil {
			log.Printf("Error executing statement: %v", err)
			return err
		}
	}

	log.Printf("Schema %s created successfully", schemaName)
	MigrationDbTenant(db)

	return
}

func DBSystemMigrate(db *sql.DB) {

	migrationsDir := path.Join("./migrations/system")
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
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
	log.Println("preparing to run migration")
	migrationsDir := path.Join("./migrations/tenants")
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error excute migration SQL for tenant, %s", err.Error()))
	} else {
		log.Println("Applied " + strconv.Itoa(n) + " migrations!")
		log.Println("Migration successfully")
	}

}
