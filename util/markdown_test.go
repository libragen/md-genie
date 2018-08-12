package util

import "testing"

func TestFetchMaoyanRedisData(t *testing.T) {
	if movies, err := FetchMaoyanRedisData(); err == nil {
		t.Log(movies)
	} else {
		t.Error(err)
	}
	RedisClient.Close()
}

func TestParseMaoyanMarkdown(t *testing.T) {
	ParseMaoyanMarkdown()
}

func TestParseMarkdownHacknews(t *testing.T) {
	ParseMarkdownHacknews()
}

func TestFetchRedisDataHackNews(t *testing.T) {
	if news, err := fetchRedisDataHackNews(); err == nil {
		t.Log(news)
	}
}

func TestParseReadme(t *testing.T) {
	ParseReadmeMarkdown()
}

func TestParseEmailContent(t *testing.T) {
	if err, emal := ParseEmailContent([]string{"logoogog"}); err == nil {
		t.Log(emal)
	} else {
		t.Error(err)
	}

}
