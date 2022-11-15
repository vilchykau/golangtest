package htmlparser

import (
	"errors"
	"testing"
)

func TestHtmlParser(t *testing.T) {
	tests := []struct {
		priceString string
		price       float64
		err         error
	}{
		{
			priceString: "123 р.",
			price:       123,
			err:         nil,
		},
		{
			priceString: "123 p.",
			price:       0,
			err:         errors.New("There is no price in the string"),
		},
	}

	kp := &KufarParser{}

	for _, testData := range tests {
		price, err := kp.parsePriceString(testData.priceString)

		if price != testData.price {
			t.Errorf("Wrong string transformation(priceString: %v)\n "+
				"price(got vs expect): %v  vs %v;\n"+
				testData.priceString, price, testData.price)
		}
		if !(err == nil && testData.err == nil) || err.Error() != testData.err.Error() {
			t.Errorf("Wrong string transformation(priceString: %v)\n "+
				"error(got vs expect): %v  vs  %v)",
				testData.priceString, err, testData.err)
		}
	}
}
