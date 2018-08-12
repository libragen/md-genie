package util

import "testing"

func TestDingLog(t *testing.T) {
	DingLog("awesome is my name!", `awesome`)
	RedisClient.Close()
}
