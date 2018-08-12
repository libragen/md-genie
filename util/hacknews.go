package util

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"time"
)

const hackNewsUrl = "http://news.ycombinator.com/news"

type NewsItem struct {
	TitleZh string `json:"titleZh"`
	TitleEn string `json:"titleEn"`
	Url     string `json:"url"`
	Date    string `json:"date"`
}

func SpiderHackNews() error {
	//stories := []item{}
	// Instantiate default collector
	doc, err := goquery.NewDocument(hackNewsUrl)
	if err != nil {
		return err
	}
	pipe := redisClient.Pipeline()
	// Find the review items
	skey := time.Now().Format("hacknews:2006-01-02")
	hkey := time.Now().Format("hacknews:2006-01")
	doc.Find("a.storylink").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		pipe.SAdd(skey, url)
		if redisClient.HGet(hkey, url).Val() == "" {
			titleEn := s.Text()
			titleZh := TranslateEn2Ch(titleEn)
			timeString := time.Now().Format("2006-01-02")
			newsItem := NewsItem{titleZh, titleEn, url, timeString}
			if bytes, err := json.Marshal(newsItem); err == nil {
				pipe.HSet(hkey, url, bytes)
			}
			time.Sleep(time.Microsecond * 100)
		}
	})
	pipe.Expire(skey, time.Hour*24)
	pipe.Expire(hkey, time.Hour*24)
	pipe.Exec()
	return nil
}
