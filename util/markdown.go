package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"
)

func FetchRedisData() ([]Movie, error) {
	skey := time.Now().Format("2006-01-02")
	ids, err := redisClient.SMembers(skey).Result()
	hkey := "maoyan_movie"
	jsonStrings, err := redisClient.HMGet(hkey, ids...).Result()

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

func ParseMarkdown() error {
	tmpl, err := template.ParseFiles("template/movies") //解析模板文件

	mdFile := fmt.Sprintf("archives/movie_%s.md", time.Now().Format("2006-01-02"))

	file, err := os.Create(mdFile)
	defer file.Close()

	movies, err := FetchRedisData()
	err = tmpl.Execute(file, movies) //执行模板的merger操作
	return err
}
