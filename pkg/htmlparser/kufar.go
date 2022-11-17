package htmlparser

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrNullPrice     = errors.New("there is no price in the string")
	ErrNegativePrice = errors.New("the price can't be negative")
)

type KufarParser struct {
	url string
}

func NewKufarParser(url string) *KufarParser {
	kf := new(KufarParser)
	kf.url = url
	return kf
}

func (kf *KufarParser) ParserPrice() (float64, error) {
	priceString, err := kf.readPage()
	if err != nil {
		return 0, err
	}

	price, err := kf.parsePriceString(priceString)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (kf *KufarParser) parsePriceString(priceString string) (float64, error) {
	re, _ := regexp.Compile(`(-?\d+) р\.`)
	res := re.FindAllStringSubmatch(priceString, 1)

	if len(res) < 1 || len(res[0]) < 2 {
		return 0, ErrNullPrice
	}

	price, err := strconv.ParseFloat(res[0][1], 64)
	if err != nil {
		return 0, err
	}

	if price < 0 {
		return 0, ErrNegativePrice
	}

	return price, nil
}

func (kf *KufarParser) readPage() (string, error) {
	resp, err := http.Get(kf.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	return doc.Find(".styles_main__PU1v4").First().Text(), nil
}
