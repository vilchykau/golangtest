package htmlparser

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type KufarParser struct {
	url string
}

func NewKufarParser(url string) *KufarParser {
	kf := new(KufarParser)
	kf.url = url
	return kf
}

func (kf *KufarParser) ParserPrice() float64 {
	re, _ := regexp.Compile(`(\d+) р\.`)
	res := re.FindAllStringSubmatch(kf.readPage(), 1)

	price, _ := strconv.ParseFloat(res[0][1], 64)

	return price
}

func (kf *KufarParser) readPage() string {
	resp, err := http.Get(kf.url)
	if err != nil {
		print(err)
		return ""
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc.Find(".styles_main__PU1v4").First().Text()
}
