package command

import (
	"flag"
	"fmt"
	"strings"

	migrate "github.com/heartz2o2o/hearts-migrate/migrate"
)

type RedoCommand struct {
}

func (c *RedoCommand) Help() string {
	helpText := `
Usage: sql-migrate redo [options] ...

  Reapply the last migration.

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  -dryrun                Don't apply migrations, just print them.

`
	return strings.TrimSpace(helpText)
}

func (c *RedoCommand) Synopsis() string {
	return "Reapply the last migration"
}

func (c *RedoCommand) Run(args []string) int {
	var dryrun bool

	cmdFlags := flag.NewFlagSet("redo", flag.ContinueOnError)
	cmdFlags.Usage = func() { fmt.Print(c.Help()) }
	cmdFlags.BoolVar(&dryrun, "dryrun", false, "Don't apply migrations, just print them.")
	ConfigFlags(cmdFlags)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	env, err := GetEnvironment()
	if err != nil {
		fmt.Printf("Could not parse config: %s", err)
		return 1
	}

	db, dialect, err := GetConnection(env)
	if err != nil {
		fmt.Print(err.Error())
		return 1
	}

	source := migrate.FileMigrationSource{
		Dir: env.Dir,
	}

	migrations, _, err := migrate.PlanMigration(db, dialect, source, migrate.Down, 1)
	if err != nil {
		fmt.Printf("Migration (redo) failed: %v", err)
		return 1
	} else if len(migrations) == 0 {
		fmt.Print("Nothing to do!")
		return 0
	}

	if dryrun {
		PrintMigration(migrations[0], migrate.Down)
		PrintMigration(migrations[0], migrate.Up)
	} else {
		_, err := migrate.ExecMax(db, dialect, source, migrate.Down, 1)
		if err != nil {
			fmt.Printf("Migration (down) failed: %s", err)
			return 1
		}

		_, err = migrate.ExecMax(db, dialect, source, migrate.Up, 1)
		if err != nil {
			fmt.Printf("Migration (up) failed: %s", err)
			return 1
		}

		fmt.Printf("Reapplied migration %s.", migrations[0].Id)
	}

	return 0
}
