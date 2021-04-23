package main

import (
	"fmt"
	"os"

	command "github.com/heartz2o2o/db-migrate/command"
	"github.com/mitchellh/cli"
)

var ui cli.Ui

func main() {
	// env := &command.Environment{
	// 	Dialect:    "mysql",
	// 	DataSource: "root:123456@tcp(localhost:3306)/bac?parseTime=true",
	// 	Dir:        "./sql"}
	// command.SetEnvironment(env)
	// command.SetIgnoreUnknown(true)
	// Upcommand := command.UpCommand{}

	// if err := Upcommand.RunProcess([]string{}); err != nil {
	// 	panic(err.Error())
	// }

	realMain()
	os.Exit(0)
}

func realMain() int {
	command.SetIgnoreUnknown(false)
	cli := &cli.CLI{
		Args: os.Args[1:],
		Commands: map[string]cli.CommandFactory{
			"up": func() (cli.Command, error) {
				return &command.UpCommand{}, nil
			},
			"down": func() (cli.Command, error) {
				return &command.DownCommand{}, nil
			},
			"redo": func() (cli.Command, error) {
				return &command.RedoCommand{}, nil
			},
			"status": func() (cli.Command, error) {
				return &command.StatusCommand{}, nil
			},
			"new": func() (cli.Command, error) {
				return &command.NewCommand{}, nil
			},
			"skip": func() (cli.Command, error) {
				return &command.SkipCommand{}, nil
			},
		},
		HelpFunc: cli.BasicHelpFunc("sql-migrate"),
		Version:  "1.0.0",
	}

	exitCode, err := cli.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
