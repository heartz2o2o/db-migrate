package command

import (
	"fmt"

	migrate "github.com/heartz2o2o/db-migrate/migrate"
)

var SetEnv *Environment

func SetEnvironment(env *Environment) {
	SetEnv = env
}

func ApplyMigrations(dir migrate.MigrationDirection, dryrun bool, limit int) error {
	var env *Environment
	yamlEnv, err := GetEnvironment()

	if err != nil && SetEnv == nil {
		return fmt.Errorf("Could not parse config: %s", err)
	} else {
		if yamlEnv != nil {
			env = yamlEnv
		} else {
			env = SetEnv
		}
	}

	db, dialect, err := GetConnection(env)
	if err != nil {
		return err
	}

	source := migrate.FileMigrationSource{
		Dir: env.Dir,
	}

	if dryrun {
		migrations, _, err := migrate.PlanMigration(db, dialect, source, dir, limit)
		if err != nil {
			return fmt.Errorf("Cannot plan migration: %s", err)
		}

		for _, m := range migrations {
			PrintMigration(m, dir)
		}
	} else {
		n, err := migrate.ExecMax(db, dialect, source, dir, limit)
		if err != nil {
			return fmt.Errorf("Migration failed: %s", err)
		}

		if n == 1 {
			fmt.Print("Applied 1 migration")
		} else {
			fmt.Printf("Applied %d migrations", n)
		}
	}

	return nil
}

func PrintMigration(m *migrate.PlannedMigration, dir migrate.MigrationDirection) {
	if dir == migrate.Up {
		fmt.Printf("==> Would apply migration %s (up)", m.Id)
		for _, q := range m.Up {
			fmt.Print(q)
		}
	} else if dir == migrate.Down {
		fmt.Printf("==> Would apply migration %s (down)", m.Id)
		for _, q := range m.Down {
			fmt.Print(q)
		}
	} else {
		panic("Not reached")
	}
}
