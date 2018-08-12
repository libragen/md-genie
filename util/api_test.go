package util

import (
	"testing"
)

func TestMovie(t *testing.T) {
	if err := FetchMaoyanApi(); err != nil {
		t.Fatal(err)
	}
	RedisClient.Close()
}
