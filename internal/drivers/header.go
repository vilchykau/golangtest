package drivers

//go:generate mockgen -destination=mocks/internal/drivers/database_mock.go -package=mocks "github.com/vilchykau/golangtest/internal/drivers" Database
//go:generate mockgen -destination=mocks/internal/drivers/tx_mock.go -package=mocks "github.com/vilchykau/golangtest/internal/drivers" Tx

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
