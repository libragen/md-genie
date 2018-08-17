package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"time"
)

type QihooSearch struct {
}

//type SearchEngine interface {
//	EngineName() string
//	SpiderRequence() time.Duration
//	Keyword() string
//	DownloadPage(urlFormat string) (io.Reader, error)
//	ParsePage(page io.Reader) ([]NewsItem, error)
//}
func (q *QihooSearch) EngineName() string {
	return "360"
}
func (q *QihooSearch) SpiderRequence() time.Duration {
	return time.Minute * 2
}
func (q *QihooSearch) Keyword() string {
	return "keyword"
}
func (q *QihooSearch) DownloadPage(urlFormat string) (io.Reader, error) {
	// Load the URL
	url := fmt.Sprintf(urlFormat, q.Keyword())
	if res, err := http.Get(url); err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		return res.Body, nil
	}
}
func (q *QihooSearch) ParsePage(page io.Reader) (list []NewsItem, err error) {
	dom, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, err
	}
	dom.Find("a.storylink").Each(func(i int, s *goquery.Selection) {

	})
	return list, err
}
