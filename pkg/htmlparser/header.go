package htmlparser

type MarketParser interface {
	ReadPrice() float64
}
