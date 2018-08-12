package util

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"time"
)

func SpiderRedditProgramming() error {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.reddit.com/r/programming/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("cookie", "edgebucket=bDuAShMKdv1UjYja8Y; loid=00000000000016lkcr.2.1490806815015.Z0FBQUFBQlpiTVBXVFFBZjVjdk1UdlhnczR1ZDJLbklscEFoSkdXeXE0MWdqZlRnQldON1hCS0w5cEllMDNLbE1HOHJqbjR3YWUzc0xCRENBZzAxamtkR3MwdnJvdU9SaU5uZV85WnVpbkxSbllKNDF4WXphbFRzTWtjNWtINkRkZnBJLVQyMERRUkU; reddaid=52HOLYHQ7L45ZJYB; rabt=; rseor3=; uapp_cookie=3%3A1527210000; USER=eyJwcmVmcyI6eyJwcm9maWxlTGF5b3V0IjoiTEFSR0UiLCJnbG9iYWxUaGVtZSI6IlJFRERJVCIsImVkaXRvck1vZGUiOiJyaWNodGV4dCIsImNvbW1lbnRNb2RlIjoicmljaHRleHQiLCJmZWF0dXJlc1ZpZXdlZEhpc3RvcnkiOnsiY29tbWVudEZvcm0iOnsibWFya2Rvd25Nb2RlTm90aWZpY2F0aW9uIjpmYWxzZX19LCJjb2xsYXBzZWRUcmF5U2VjdGlvbnMiOnsiZmF2b3JpdGVzIjpmYWxzZSwibXVsdGlzIjpmYWxzZSwibW9kZXJhdGluZyI6ZmFsc2UsInN1YnNjcmlwdGlvbnMiOmZhbHNlLCJwcm9maWxlcyI6ZmFsc2V9LCJuaWdodG1vZGUiOmZhbHNlLCJzdWJzY3JpcHRpb25zUGlubmVkIjpmYWxzZX0sImxhbmd1YWdlIjoiZW4ifQ==; recent_srs=t5_2rc7j%2C; trytvorg_recent_srs=t5_2qqjc%2Ct5_2fwo%2Ct5_2rc7j%2Ct5_2w7ch%2Ct5_2qyw8%2Ct5_2qh33%2Ct5_3m0tc%2Ct5_2qh0s%2Ct5_2qq6z%2Ct5_2tycb; reddit_session=71550027%2C2018-07-20T17%3A11%3A48%2C53242cc87693973843977bd38387c8b426923d99; pc=xp; _recent_srs=t5_2fwo%2Ct5_3mmij%2Ct5_2qnlf%2Ct5_2qmfx%2Ct5_2qt55%2Ct5_2qh0u%2Ct5_2rc7j%2Ct5_2qh72%2Ct5_2qh03%2Ct5_2qh3s; listingsignupbar_dismiss=1; initref=google.com.tw; session_tracker=NQ7ovVExLDtAkYsClZ.0.1532133588474.Z0FBQUFBQmJVb0RVUEx1TnRCOU1TczVtbTVBa2FxalRGcmVvOVUtN2NSRXlUYUk1MHV3MzUyWElra2FzZzhzdi1MdXVoaHhJa2IycDdlTHFKa2NHd241MnZ4cC1zUXRTZkpoSWg5SFpfOGs3Mk1QSTFZWm5oMFAwbmxIcnFteE94SWtPYlExOU5yTVM")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(doc.Contents().Text())
	pipe := RedisClient.Pipeline()
	// Find the review items
	skey := time.Now().Format("redditnews-2006-01-02")
	hkey := "redditnews"
	doc.Find("#siteTable > div.link > div.entry.unvoted > div.top-matter > p.title > a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Print(url)
		pipe.SAdd(skey, url)
		if RedisClient.HGet(hkey, url).Val() == "" {
			titleEn := s.Text()
			titleZh := TranslateEn2Ch(titleEn)
			timeString := time.Now().Format("2006-01-02")
			newsItem := NewsItem{titleZh, titleEn, url, timeString}
			if bytes, err := json.Marshal(newsItem); err == nil {
				pipe.HSet(hkey, url, bytes)
			}
			time.Sleep(time.Microsecond * 100)
		}
	})
	pipe.Expire(skey, time.Hour*12)
	pipe.Exec()
	return nil
}
