package runner

import (
	"bytes"
	"io"
	"net/http"
	"sync"
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

func Run(cfg RequestConfig) []Result {
	client := &http.Client{
		Timeout: cfg.Timeout,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: cfg.Concurrency,
		},
	};

	jobs := make(chan struct{}, cfg.TotalReqs);  // to hold all the jobs
	results := make(chan Result, cfg.TotalReqs);  // to hold all the results

	for i := 0;i < cfg.TotalReqs;i++ {
		jobs <- struct{}{};
	}

	close(jobs);

	// workers now running as jobs are done added all
	var wg sync.WaitGroup;
	for i := 0 ;i < cfg.Concurrency;i++ {
		wg.Add(1);
		go func() {
			defer wg.Done()
			for range jobs {
				res := doRequest(client, cfg);
				results <- res;
			}
		}()
	}

	go func ()  {
		wg.Wait();
		close(results);
	}()

	var resList []Result
	for r := range results {
		resList = append(resList, r)
	}
	return resList
}