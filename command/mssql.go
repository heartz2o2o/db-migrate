// +build go1.3

package command

import (
	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/gorp.v1"
)

func init() {
	dialects["mssql"] = gorp.SqlServerDialect{}
}
