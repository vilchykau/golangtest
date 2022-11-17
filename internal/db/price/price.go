package price

import (
	"github.com/jmoiron/sqlx"
	"github.com/vilchykau/golangtest/internal/comerror"
	"github.com/vilchykau/golangtest/internal/drivers"
)

func NewPrice(price float64, url string) *Price {
	p := new(Price)
	p.Price = &price
	p.Url = &url
	return p
}

func (p *Price) InsertPrice(db drivers.Database) error {
	var res *sqlx.Rows
	var err error

	tx := db.MustBegin()

	if p.PriceID == nil {
		res, err = tx.NamedQuery(
			`INSERT INTO MARKET.T_PRICE (URL, PRICE)
			 VALUES(:URL, :PRICE)
			 RETURNING *`, p)
	} else {
		res, err = tx.NamedQuery(
			`INSERT INTO MARKET.T_PRICE(PRICE_ID, URL, PRICE)
			 VALUES(:PRICE_ID, :URL, :PRICE)
			 RETURNING *`, p)
	}
	tx.Commit()

	if err != nil {
		return err
	}
	defer res.Close()

	if !res.Next() {
		return comerror.ErrUnknownError
	}

	return nil
}
