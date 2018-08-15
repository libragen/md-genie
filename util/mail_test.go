package util

import "testing"

func TestSendMsgToEmail(t *testing.T) {
	if err := SendMsgToEmail("test title", "test body"); err != nil {
		t.Log(err)
	}

}
