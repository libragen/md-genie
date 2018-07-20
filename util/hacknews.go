package util

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"time"
)

const hackNewsUrl = "http://news.ycombinator.com/news"

const (
	newsFeed = "news_feed"
	newsPost = "news_post"
	newsList = "hack_news"
)

type HacknewsItem struct {
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
		return nil
	}
	pipe := redisClient.Pipeline()
	// Find the review items
	skey := time.Now().Format("hacknews-2006-01-02")
	doc.Find("a.storylink").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		//url = strings.TrimSpace(url)
		pipe.SAdd(skey, url)
		if redisClient.HGet(newsList, url).Val() == "" {
			titleEn := s.Text()
			titleZh := TranslateEn2Ch(titleEn)
			timeString := time.Now().Format("2006-01-02")
			newsItem := HacknewsItem{titleZh, titleEn, url, timeString}
			if bytes, err := json.Marshal(newsItem); err == nil {
				pipe.HSet(newsList, url, bytes)
			}
			time.Sleep(time.Microsecond * 100)
		}
	})
	pipe.Expire(skey, time.Hour*12)
	pipe.Exec()
	return nil
}
