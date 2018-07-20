package util

import "testing"

func TestFetchRedisData(t *testing.T) {
	if movies, err := FetchRedisData(); err == nil {
		t.Log(movies)
	} else {
		t.Error(err)
	}
}

func TestParseMarkdown(t *testing.T) {
	ParseMarkdown()
}
