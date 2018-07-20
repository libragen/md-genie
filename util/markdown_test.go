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

}

func TestFetchRedisDataHackNews(t *testing.T) {
	if news,err := FetchRedisDataHackNews();err ==nil {
		t.Log(news)
	}
	redisClient.Close()

}
