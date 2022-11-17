package drivers

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ErrNotInitDB = errors.New("the db connection hasn't been initialized")
)

type PgConfig struct {
	DnsString string `env:"POSTGRES_DSN" validate:"required"`
}

type PgDriver struct {
	config PgConfig
	db     *sqlx.DB
}

func NewPgDriver(config PgConfig) Database {
	return &PgDriver{
		config: config,
	}
}

func (d *PgDriver) Connect() error {
	db, err := sqlx.Connect("postgres", d.config.DnsString)
	if err != nil {
		return err
	}
	d.db = db

	return nil
}

func (d *PgDriver) Close() error {
	err := d.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (d *PgDriver) GetDb() *sqlx.DB {
	return d.db
}

func (d *PgDriver) MustBegin() Tx {
	return d.db.MustBegin()
}
