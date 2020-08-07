// +build oracle

package command

import (
	migrate "github.com/heartz2o2o/db-migrate/migrate"
	_ "github.com/mattn/go-oci8"
)

func init() {
	dialects["oci8"] = migrate.OracleDialect{}
}
