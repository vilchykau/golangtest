package controller

import (
	"errors"
	"net/http"

	"github.com/vilchykau/golangtest/internal/store"
	"github.com/vilchykau/golangtest/pkg/db"
	"github.com/vilchykau/golangtest/pkg/htmlparser"
)

func AddSubcription(email string, url string) (int, string) {
	db := db.GetDatabase()

	parser := htmlparser.NewKufarParser(url)
	priceV, err := parser.ParserPrice()
	if err != nil {
		return http.StatusInternalServerError, "Wrong address"
	}

	err = store.NewPrice(priceV, url).InsertPrice(db)
	if !errors.Is(err, store.ErrPriceAlreadyExists) {
		return http.StatusInternalServerError, "Internal Error!"
	}

	err = store.NewSubcription(email, url).Insert(db)
	if err == nil {
		return http.StatusOK, "OK"
	} else if errors.Is(err, store.ErrSubcriptionAlreadyExists) {
		return http.StatusOK, "Already"
	} else {
		return http.StatusInternalServerError, "Internal Error!"
	}
}
