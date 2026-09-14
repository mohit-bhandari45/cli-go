package runner

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func buildRequest(cfg RequestConfig) (*http.Request, error) {
	var bodyReader io.Reader
	if len(cfg.Body) > 0 {
		bodyReader = bytes.NewReader(cfg.Body)
	}

	req, err := http.NewRequest(cfg.Method, cfg.URL, bodyReader)
	if err != nil {
		return nil, err
	}

	for key, values := range cfg.Headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}

func doRequest(client *http.Client, cfg RequestConfig) Result {
	req, err := buildRequest(cfg)
	if err != nil {
		return Result{Error: err.Error()}
	}

	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		return Result{
			Duration: duration,
			Error:    err.Error(),
		}
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	return Result{
		StatusCode: resp.StatusCode,
		Duration:   duration,
		BytesRead:  int64(len(bodyBytes)),
	}
}
