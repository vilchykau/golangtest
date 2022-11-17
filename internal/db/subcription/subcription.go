package subcription

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/vilchykau/golangtest/internal/comerror"
)

func NewSubcription(email string, url string) *Subcription {
	return &Subcription{
		Email: &email,
		Url:   &url,
	}
}

func (s *Subcription) Insert(ctx *gin.Context) error {
	//var res sql.Result
	var err error

	var dbRaw, ok = ctx.Get("DB")
	if !ok {
		return comerror.ErrDatabaseNotInit
	}

	db := dbRaw.(*sqlx.DB)
	tx := db.MustBegin()

	_, err = tx.NamedExec(`
		INSERT INTO MARKET.T_SUBCRIPTION(EMAIL, PRICE_ID)
		SELECT :EMAIL, PRICE_ID
		FROM MARKET.T_PRICE
		WHERE URL = :URL
	`, s)
	tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
