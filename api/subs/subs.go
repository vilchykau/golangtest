package subs

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/vilchykau/golangtest/internal/db/price"
	"github.com/vilchykau/golangtest/internal/db/subcription"
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

	parser := htmlparser.NewKufarParser(body.Url)
	priceV, _ := parser.ParserPrice()
	price.NewPrice(priceV, body.Url).InsertPrice(g)

	subcription.NewSubcription(body.Email, body.Url).Insert(g)

	g.JSON(http.StatusOK, body.Email)
}

func InitGroup(g *gin.RouterGroup) {
	g.POST("/addsubs", AddSubcription)
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
