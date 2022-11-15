package htmlparser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
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
		{
			priceString: "123 %",
			price:       0,
			err:         errors.New("There is no price in the string"),
		},
		{
			priceString: "-1 р.",
			price:       0,
			err:         errors.New("The price can't be negative"),
		},
	}

	kp := &KufarParser{}

	for _, testData := range tests {
		price, err := kp.parsePriceString(testData.priceString)
		t.Run("Kufar regex parser", func(t *testing.T) {
			require.Equal(t, price, testData.price)
			if testData.err == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, testData.err.Error())
			}
		})
	}
}
