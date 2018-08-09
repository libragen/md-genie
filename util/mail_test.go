package util

import "testing"

func TestSendMsgToEmail(t *testing.T) {
	if err := SendMsgToEmail("test title","test body","erikchau@me.com");err != nil {
		t.Log(err)
	}

}
