package runner

import (
	"net/http"
	"time"
)

// request config to store request config
type RequestConfig struct {
	URL 			string
	Method			string
	Headers 		http.Header
	Body			[]byte
	Timeout			time.Duration
	TotalReqs		int
	Concurrency		int
}