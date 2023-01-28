package command

import (
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/apsyadira-jubelio/migration-tools-database/driver"
	"github.com/mitchellh/cli"
)

type MigrateAllTenants struct {
	Ui cli.Ui
}

type Tenants struct {
	Email      string
	SchemaName string
	Host       string
	Port       string
}

func (c *MigrateAllTenants) Help() string {
	helpText := `
Usage: jb-chat-migrate all-tenants
  Create schema database and run migration for all tenants
`
	return strings.TrimSpace(helpText)
}

func (c *MigrateAllTenants) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("new", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	systemDb := driver.PostgreDbClient(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var tenants = []Tenants{}

	sqlString := "select email, schema_name, hostname from users where is_owner is true"
	rows, err := systemDb.Query(sqlString)

	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	for rows.Next() {
		var tenant Tenants
		err := rows.Scan(&tenant.Email, &tenant.SchemaName, &tenant.Host)
		if err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
		tenants = append(tenants, tenant)
	}

	var wg sync.WaitGroup
	wg.Add(len(tenants))
	for _, data := range tenants {
		tenantDb := driver.PostgreDbClient(data.Host, os.Getenv("DB_TENANT_PORT"), os.Getenv("DB_TENANT_USER"), os.Getenv("DB_TENANT_PASSWORD"), os.Getenv("DB_TENANT_NAME"))
		go driver.CreatePostgreSchema(tenantDb, data.SchemaName, &wg)
	}
	wg.Wait()
	return 0
}

func (c *MigrateAllTenants) Synopsis() string {
	return "Run create schema and migration for specific tenant"
}
