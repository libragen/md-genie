package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"
	"time"
)

func FetchMaoyanRedisData() ([]Movie, error) {
	skey := time.Now().Format("maoyan:2006-01-02")
	ids, err := RedisClient.SMembers(skey).Result()
	hkey := "maoyan:movie"
	jsonStrings, err := RedisClient.HMGet(hkey, ids...).Result()

	movies := []Movie{}
	for _, item := range jsonStrings {
		if string, ok := item.(string); ok {
			movie := Movie{}
			json.Unmarshal([]byte(string), &movie)
			movies = append(movies, movie)
		}
	}

	return movies, err
}

func ParseMaoyanMarkdown() error {
	tmpl, err := template.ParseFiles("template/movies") //解析模板文件

	mdFile := fmt.Sprintf("archives/movie_%s.md", time.Now().Format("2006-01-02"))

	file, err := os.Create(mdFile)
	defer file.Close()

	movies, err := FetchMaoyanRedisData()
	err = tmpl.Execute(file, movies) //执行模板的merger操作
	return err
}

func fetchRedisDataHackNews() ([]NewsItem, error) {
	skey := time.Now().Format("hacknews:2006-01-02")
	urls, err := RedisClient.SMembers(skey).Result()
	if err != nil || len(urls) == 0 {
		return nil, errors.New("没有新闻 smembers")
	}
	hkey := time.Now().Format("hacknews:2006-01")

	jsonStrings, err := RedisClient.HMGet(hkey, urls...).Result()
	if err != nil || len(jsonStrings) == 0 {
		return nil, errors.New("没有新闻 hmget")
	}
	newsItems := []NewsItem{}
	for _, item := range jsonStrings {
		if string, ok := item.(string); ok {
			items := NewsItem{}
			json.Unmarshal([]byte(string), &items)
			newsItems = append(newsItems, items)
		}
	}

	return newsItems, err
}
func ParseMarkdownHacknews() error {
	tmpl, err := template.ParseFiles("template/hacknews") //解析模板文件

	mdFile := fmt.Sprintf("archives/hacknews_%s.md", time.Now().Format("2006-01-02"))

	file, err := os.Create(mdFile)
	defer file.Close()

	newsItems, err := fetchRedisDataHackNews()
	err = tmpl.Execute(file, newsItems) //执行模板的merger操作
	return err
}

func ParseEmailContent(logs []string) (error, string) {
	tmpl, err := template.ParseFiles("template/mailNews.html") //解析模板文件
	if err != nil {
		return err, ""
	}

	newsItems, err := fetchRedisDataHackNews()

	var buf = new(bytes.Buffer)
	data := struct {
		List []NewsItem
		Logs []string
	}{newsItems, logs}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return err, ""
	}
	return nil, buf.String()
}

func ParseReadmeMarkdown() error {
	tmpl, err := template.ParseFiles("template/home") //解析模板文件
	file, err := os.Create("README.md")
	defer file.Close()

	items, err := LsArchivesMdFiles("archives")
	err = tmpl.Execute(file, items) //执行模板的merger操作
	return err
}
