package command

import (
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/apsyadira-jubelio/migration-tools-database/driver"
	"github.com/mitchellh/cli"
)

type MigrateDbSystem struct {
	Ui cli.Ui
}

func (c *MigrateDbSystem) Help() string {
	helpText := `
Usage: jb-chat-migrate db-system
  Run migration for database system
`
	return strings.TrimSpace(helpText)
}

func (c *MigrateDbSystem) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("new", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	systemDb := driver.PostgreDbClient(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		driver.DBSystemMigrate(systemDb)
		defer wg.Done()
	}()
	wg.Wait()

	return 0
}

func (c *MigrateDbSystem) Synopsis() string {
	return "Run all migration for database system"
}
