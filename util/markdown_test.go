package util

import "testing"

func TestFetchRedisData(t *testing.T) {
	if movies, err := FetchRedisData(); err == nil {
		t.Log(movies)
	} else {
		t.Error(err)
	}
	redisClient.Close()
}

func TestParseMarkdown(t *testing.T) {
	ParseMarkdown()
	redisClient.Close()
}

func TestParseMarkdownHacknews(t *testing.T) {
	ParseMarkdownHacknews()
}

func TestFetchRedisDataHackNews(t *testing.T) {
	if news, err := fetchRedisDataHackNews(); err == nil {
		t.Log(news)
	}
}
