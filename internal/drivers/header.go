package drivers

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	Connect() error
	Close() error
	MustBegin() Tx
}

type Tx interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Commit() error
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}
