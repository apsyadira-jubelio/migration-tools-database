package main

import (
	"fmt"
	"os"

	"github.com/apsyadira-jubelio/migration-tools-database/command"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mitchellh/cli"
)

func main() {
	godotenv.Load()
	os.Exit(realMain())
}

func realMain() int {

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	cli := &cli.CLI{
		Args: os.Args[1:],
		Commands: map[string]cli.CommandFactory{
			"tenant": func() (cli.Command, error) {
				return &command.MigrateSpecificTenant{Ui: ui}, nil
			},
			"all-tenant": func() (cli.Command, error) {
				return &command.MigrateAllTenants{Ui: ui}, nil
			},
			"db-system": func() (cli.Command, error) {
				return &command.MigrateDbSystem{Ui: ui}, nil
			},
		},
		HelpFunc: cli.BasicHelpFunc("jb-chat-migrate"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
