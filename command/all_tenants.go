package command

import (
	"flag"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/apsyadira-jubelio/migration-tools-database/config"
	"github.com/apsyadira-jubelio/migration-tools-database/driver"
	"github.com/apsyadira-jubelio/migration-tools-database/utils"
	"github.com/mitchellh/cli"
)

type MigrateAllTenants struct {
	Ui *cli.BasicUi
}

type Tenants struct {
	Email      string
	SchemaName string
	Host       string
	Port       string
	DBName     string
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

	systemDb := driver.PostgreDbClient(config.Config.System.Datasource)

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

	tw := tabwriter.NewWriter(c.Ui.Writer, 0, 0, 3, ' ', 0)
	defer tw.Flush()
	fmt.Fprintln(tw, strings.Repeat("-", utils.GetTerminalWidth()))
	for _, data := range tenants {
		fmt.Println("Hostname", data.Host, "Schemaname", data.SchemaName)
		fmt.Fprintln(tw, strings.Repeat("-", utils.GetTerminalWidth()))
		tenantDb := driver.PostgreDbClient(config.Config.Tenant.Datasource)
		driver.CreatePostgreSchema(tenantDb, data.SchemaName)
		fmt.Fprintln(tw, strings.Repeat("-", utils.GetTerminalWidth()))
	}
	return 0
}

func (c *MigrateAllTenants) Synopsis() string {
	return "Run create schema and migration for specific tenant"
}
