package engine

import (
	"github.com/sirupsen/logrus"
	"lh-example/lh-crawler/single/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			logrus.Errorln("worker失败", err)
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			logrus.Infof("获取到item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	logrus.Infoln("读取URL", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		logrus.Error("fetch失败", err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil

}
