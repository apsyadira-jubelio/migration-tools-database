package command

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/apsyadira-jubelio/migration-tools-database/config"
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

	fmt.Println(config.Config.System.Datasource)
	systemDb := driver.PostgreDbClient(config.Config.System.Datasource)

	defer systemDb.Close()

	if !driver.EnsureDbConnection(systemDb, 3) {
		c.Ui.Error("Failed to reconnect to system database")
		return 1
	}

	var tenantData Tenants
	err := systemDb.QueryRow("select schema_name, hostname from users where schema_name = $1", cmdFlags.Arg(0)).Scan(&tenantData.SchemaName, &tenantData.Host)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Tenant %s not found", cmdFlags.Arg(0))
			return 1
		} else {
			log.Fatalf(err.Error())
			return 1
		}

	}

	tenantDb := driver.PostgreDbClient(config.Config.Tenant.Datasource)
	defer tenantDb.Close()

	if !driver.EnsureDbConnection(tenantDb, 3) {
		c.Ui.Error("Failed to reconnect to tenant database")
		return 1
	}

	if err := driver.CreatePostgreSchema(tenantDb, tenantData.SchemaName); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func (c *MigrateSpecificTenant) Synopsis() string {
	return "Run create schema and migration for specific tenant"
}
