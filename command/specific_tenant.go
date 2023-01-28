package command

import (
	"database/sql"
	"errors"
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/apsyadira-jubelio/migration-tools-database/driver"
	"github.com/mitchellh/cli"
)

type MigrateSpecificTenant struct {
	Ui cli.Ui
}

func (c *MigrateSpecificTenant) Help() string {
	helpText := `
Usage: jb-chat-migrate tenant schemaname
  Create schema database and run migration for specific tenant

Options:
	schemaname                   The schemaname database of tenants
`
	return strings.TrimSpace(helpText)
}

func (c *MigrateSpecificTenant) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("new", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	if len(args) < 1 {
		err := errors.New("a schemaname for the migration is needed")
		c.Ui.Error(err.Error())
		return 1
	}

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	systemDb := driver.PostgreDbClient(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var tenantData Tenants
	err := systemDb.QueryRow("select schema_name, hostname from users where schema_name = $1", cmdFlags.Arg(0)).Scan(&tenantData.SchemaName, &tenantData.Host)

	if err != nil {
		if err == sql.ErrNoRows {
			c.Ui.Error("Tenant not found")
			return 1
		} else {
			c.Ui.Error(err.Error())
			return 1
		}

	}

	tenantDb := driver.PostgreDbClient(tenantData.Host, os.Getenv("DB_TENANT_PORT"), os.Getenv("DB_TENANT_USER"), os.Getenv("DB_TENANT_PASSWORD"), os.Getenv("DB_TENANT_NAME"))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		driver.CreatePostgreSchema(tenantDb, tenantData.SchemaName, &wg)
	}()
	wg.Wait()

	return 0
}

func (c *MigrateSpecificTenant) Synopsis() string {
	return "Run create schema and migration for specific tenant"
}
