package lsapi

import (
	"strings"

	"bitbucket.org/okestrolab/baton-om-sdk/btoutil/restapi"
)

type Logstash struct {
	env struct {
		urls []string
	}

	_url_idx int
}

func (r *Logstash) Init(urls string) {
	r.env.urls = strings.Split(strings.ReplaceAll(strings.Trim(urls, "{[]}"), " ", ""), ",")
	r._url_idx = 0
}

func (r *Logstash) Send(msg interface{}) (err error) {
	for i := 0; i < len(r.env.urls); i++ {
		sndIdx := (r._url_idx + i) % len(r.env.urls)
		_, _, err = restapi.RqstPost(r.env.urls[sndIdx], nil, msg, nil)
		if err != nil {
			r._url_idx = sndIdx
			break
		}
	}

	return err
}
