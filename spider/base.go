package spider

import (
	"github.com/dejavuzhou/md-genie/config"
	"github.com/go-redis/redis"
	"time"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"errors"
)

type SubscribeJob struct {
	SpiderFrequency      time.Duration
	SubscribeFrequency      time.Duration
	Keyword        string
	SelectedEngine []string
}

type NewsItem struct {
	Title       string    `json:title`
	Description string    `json:description`
	Source      string    `json:source`
	Url         string    `json:url`
	CreatedAt   time.Time `json:created_at`
}
//spider engine 需要实现是的方法
type SearchEngine interface {
	EngineName() string
	UrlFormat() string
	Keyword() string
	ParsePage(document *goquery.Document) ([]NewsItem, error)
}

var spider *Spider
//初始化redis
func init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD, // no password set
		DB:       config.REDIS_DB_IDX,   // use default DB
	})
	if pong, err := redisClient.Ping().Result(); err != nil || pong !="PONG" {
		log.Fatalln("initializing spider has failled for the result of pinging redis is not PONG or redis configuration is wrong")
	}
	spider = &Spider{map[string]SearchEngine{}, redisClient}
}

type Spider struct {
	Engines map[string]SearchEngine
	redis   *redis.Client
}
//注册spider
func RegisterEngine(engine SearchEngine){
	engineName := engine.EngineName()
	spider.Engines[engineName] = engine
}
func (s *Spider) Work(job *SubscribeJob) error {
	return nil
}
//下载搜索关键字的html
func (s *Spider) downloadHtml(urlFormat, keyword string) (*goquery.Document,error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(urlFormat,keyword), nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("cookie", config.HTTP_COOKIE)
	req.Header.Set("User-Agent",config.HTTP_USER_AGENT)
	res, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("the get request's response code is not 200")
	}
	defer res.Body.Close()
	return  goquery.NewDocumentFromReader(res.Body)
}
//抓取新闻
func (s *Spider) FetchNews(engineName, keyword string)error  {
	//1 find the engine
	engine,ok := s.Engines[engineName]
	if ok== false {
		return errors.New("can not find registered search engine of key:"+keyword)
	}
	dom,err := s.downloadHtml(engine.UrlFormat(),engine.Keyword())
	if err != nil {
		return err
	}
	newsList,err := engine.ParsePage(dom)
	if err != nil {
		return err
	}
	return s.storeData(keyword,newsList)
}

func (s *Spider)storeData(keyword string,newsList []NewsItem) error  {
	return nil
}
