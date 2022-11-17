package subcription

import (
	"errors"

	"github.com/vilchykau/golangtest/internal/drivers"
)

var (
	ErrSubcriptionAlreadyExists = errors.New("pq: duplicate key value violates unique constraint \"t_subcription_email_price_id_key\"")
)

func NewSubcription(email string, url string) *Subcription {
	return &Subcription{
		Email: &email,
		Url:   &url,
	}
}

func (s *Subcription) Insert(db drivers.Database) error {
	//var res sql.Result
	var err error

	tx := db.MustBegin()

	_, err = tx.NamedExec(`
		INSERT INTO MARKET.T_SUBCRIPTION(EMAIL, PRICE_ID)
		SELECT :EMAIL, PRICE_ID
		FROM MARKET.T_PRICE
		WHERE URL = :URL
	`, s)
	tx.Commit()

	if err != nil && err.Error() == ErrSubcriptionAlreadyExists.Error() {
		return ErrSubcriptionAlreadyExists
	} else if err != nil {
		return err
	}

	return nil
}
