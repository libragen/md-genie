package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

/*
	有道智云
	文档页面  http://ai.youdao.com/docs/doc-trans-api.s#p01
*/
const appId = "11dfdb97abcbaed3"
const appSecret = "onyaSnhdcYIZlDHQLWW6No0stUbs5Cxy"
const host = "http://openapi.youdao.com/api"


func TranslateCh2En(text string) string{
	res := translateChinese2English(text,"zh-CHS","EN")
	if len(res.Translation) > 0 {
		return res.Translation[0]
	}else {
		return ""
	}
}
func TranslateEn2Ch(text string) string {
	res := translateChinese2English(text,"EN","zh-CHS")
	if len(res.Translation) > 0 {
		return res.Translation[0]
	}else {
		return ""
	}
}

func translateChinese2English(text,from,to string) (obj *responseStruct) {
	salt := strconv.Itoa(rand.Intn(999))
	sign := generateSign(text, salt)
	data := url.Values{
		"q":      {text},
		"to":   {to},
		"from":     {from},
		"appKey": {appId},
		"salt":   {salt},
		"sign":   {sign},
		"ext":    {"mp3"},
		"voice":  {"0"},
	}
	resp, err := http.PostForm(host, data)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Print(err)
	}
	obj = &responseStruct{}
	json.NewDecoder(resp.Body).Decode(obj)
	return obj
}

type responseStruct struct {
	TSpeakURL   string   `json:"tSpeakUrl"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	ErrorCode   string   `json:"errorCode"`
	Dict        struct {
		URL string `json:"url"`
	} `json:"dict"`
	Webdict struct {
		URL string `json:"url"`
	} `json:"webdict"`
	L        string `json:"l"`
	SpeakURL string `json:"speakUrl"`
}

func generateSign(q, salt string) string {
	temp := appId + q + salt + appSecret
	h := md5.New()
	io.WriteString(h, temp)
	return hex.EncodeToString(h.Sum(nil))
}
