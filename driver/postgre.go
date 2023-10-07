package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

func PostgreDbClient(dataSource string) *sql.DB {
	db, err := sql.Open("postgres", dataSource)

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

func CreatePostgreSchema(db *sql.DB, schemaName string) (err error) {
	log.Printf("Run migrations & create schema of %s", schemaName)

	schemaSql := []string{
		fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s AUTHORIZATION %s;`, schemaName, os.Getenv("DATABASE_TENANT_USERNAME")),
		fmt.Sprintf("SET search_path TO %s, pg_catalog;", schemaName),
	}

	for _, stringCmd := range schemaSql {
		stmt, err := db.Prepare(stringCmd)

		if err != nil {
			log.Fatalf("Error preparing statement: %v", err)
			return err
		}

		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			log.Fatalf("Error executing statement: %v", err)
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
