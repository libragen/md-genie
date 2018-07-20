package util

import (
	"bytes"
	"fmt"
	"net/http"
)

const DingApiUrl = "https://oapi.dingtalk.com/robot/send?access_token=e28e8b2efdd05a9954f888ab16b2e059706628f85590e269cae996eae7fbbf8f"

func DingLog(content, title string) error {
	if content != "" {
		format := `
		{
			"msgtype": "markdown",
			"markdown": {
				"title":"%s",
				"text": "%s"
			}
		}`
		body := fmt.Sprintf(format, title, content)
		jsonValue := []byte(body)
		_, err := http.Post(DingApiUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return err
		}
	}
	return nil
}
