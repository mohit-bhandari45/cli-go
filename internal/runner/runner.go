package runner

import (
	"bytes"
	"io"
	"net/http"
)

func buildRequest(cfg RequestConfig) (*http.Request, error) {
	var bodyReader io.Reader;
	if len(cfg.Body) > 0 {
		bodyReader = bytes.NewReader(cfg.Body);
	}

	req, err := http.NewRequest(cfg.Method, cfg.URL, bodyReader);
	if err != nil {
		return nil, err;
	}

	for key, values := range cfg.Headers {
		for _, value := range values {
			req.Header.Add(key, value);
		}
	}

	return req, nil;
}
