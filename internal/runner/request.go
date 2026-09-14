package runner

import (
	"net/http"
	"time"
)

type RequestConfig struct {
	URL 			string
	Method			string
	Headers 		http.Header
	Body			[]byte
	Timeout			time.Duration
	TotalReqs		int
	Concurrency		int
}