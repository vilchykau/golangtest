package price

type Price struct {
	PriceID *int64   `db:"PRICE_ID" json:"price_id"`
	Url     *string  `db:"URL" json:"url"`
	Price   *float64 `db:"PRICE" json:"price"`
}
