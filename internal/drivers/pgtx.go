package drivers

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PgTx struct {
	tx *sqlx.Tx
}

func NewPgTx(tx *sqlx.Tx) *PgTx {
	return &PgTx{tx: tx}
}

func (pt *PgTx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return pt.tx.NamedExec(query, arg)
}

func (pt *PgTx) Commit() error {
	return pt.tx.Commit()
}

func (pt *PgTx) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return pt.tx.NamedQuery(query, arg)
}
