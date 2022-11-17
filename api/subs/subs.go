package subs

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/vilchykau/golangtest/internal/db/price"
	"github.com/vilchykau/golangtest/internal/db/subcription"
	"github.com/vilchykau/golangtest/internal/drivers"
	"github.com/vilchykau/golangtest/pkg/htmlparser"
)

type AddSubcriptionBody struct {
	Email string `json:"email" validate:"required" example:"Subcriber@email.com"`
	Url   string `json:"url" validate:"required" example:"https:https://kufar.by"`
}

// AddSubcription godoc
// @Summary add Subcription to email.
// @Description add Subcription to email.
// @Tags api/v1/subs/
// @Accept */*
// @Param Body body AddSubcriptionBody true "The body to create a thing"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/subs/addsubs [post]
func AddSubcription(g *gin.Context) {
	var body AddSubcriptionBody
	g.BindJSON(&body)

	if !validateEmail(body.Email) {
		g.JSON(http.StatusBadRequest, "Wrong email!")
		return
	}

	db, _ := drivers.UnpackDatabase(g)

	parser := htmlparser.NewKufarParser(body.Url)
	priceV, err := parser.ParserPrice()
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Wrong address")
		return
	}

	err = price.NewPrice(priceV, body.Url).InsertPrice(db)
	if !errors.Is(err, price.ErrPriceAlreadyExists) {
		g.JSON(http.StatusInternalServerError, "Internal Error!")
		return
	}

	err = subcription.NewSubcription(body.Email, body.Url).Insert(db)
	if err == nil {
		g.JSON(http.StatusOK, "OK")
	} else if errors.Is(err, subcription.ErrSubcriptionAlreadyExists) {
		g.JSON(http.StatusOK, "Already")
	} else {
		g.JSON(http.StatusInternalServerError, "Internal Error!")
	}
}

func InitGroup(g *gin.RouterGroup) {
	g.POST("/addsubs", AddSubcription)
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
