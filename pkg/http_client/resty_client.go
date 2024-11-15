package http_client

import (
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	DefaultRetryTime = 3
	DefaultTimeout   = 2
)

func NewRestyClient(retryTime, timeout int, ENV string) *resty.Client {
	client := resty.New()

	client.SetRetryCount(DefaultRetryTime).SetTimeout(time.Duration(DefaultTimeout) * time.Second)
	if retryTime != 0 {
		client.SetRetryCount(retryTime)
	}
	if timeout != 0 {
		client.SetTimeout(time.Duration(timeout) * time.Second)
	}

	client.AddRetryCondition(customRetryCondition)

	client.SetDebugBodyLimit(200)
	if ENV == "dev" {
		client.SetDebug(true)
	}
	return client
}

func customRetryCondition(r *resty.Response, err error) bool {
	return err != nil || (r.StatusCode() >= 500 && r.StatusCode() != 503)
}
