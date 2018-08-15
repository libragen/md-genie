package util

import "testing"

func TestSendMsgToEmail(t *testing.T) {
	_, mailBody := ParseEmailContent([]string{"adfsafdasf", "sdadsA"})

	if err := SendMsgToEmail("test title", mailBody); err != nil {
		t.Log(err)
	}

}
