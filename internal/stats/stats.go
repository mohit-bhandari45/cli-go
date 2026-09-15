package stats

import (
	"sort"
	"time"

	"github.com/mohit-bhandari45/gurl/internal/runner"
)

type Summary struct {
	TotalRequests   int            `json:"total_requests"`
	SuccessRequests int            `json:"success_requests"`
	FailedRequests  int            `json:"failed_requests"`
	TotalDuration   time.Duration  `json:"total_duration_ns"`
	RPS             float64        `json:"requests_per_second"`
	MinDuration     time.Duration  `json:"min_duration_ns"`
	MaxDuration     time.Duration  `json:"max_duration_ns"`
	AvgDuration     time.Duration  `json:"avg_duration_ns"`
	P50             time.Duration  `json:"p50_duration_ns"`
	P95             time.Duration  `json:"p95_duration_ns"`
	P99             time.Duration  `json:"p99_duration_ns"`
	StatusCodes     map[int]int    `json:"status_codes"`
	Errors          map[string]int `json:"errors,omitempty"`
}

func calculatePercentile(durations []time.Duration, p float64) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	index := int(float64(len(durations)) * p)
	if index > len(durations) {
		index = len(durations) - 1
	}

	return durations[index]
}

func Calculate(results []runner.Result, totalDuration time.Duration) Summary {
	sum := Summary{
		TotalRequests: len(results),
		TotalDuration: totalDuration,
		StatusCodes:   make(map[int]int),
		Errors:        make(map[string]int),
	}

	var durations []time.Duration;
	var totalDurationSum time.Duration;

	for _, res := range results {
		if res.Error != "" {
			sum.Errors[res.Error]++
			sum.FailedRequests++
		} else {
			if res.StatusCode >= 200 && res.StatusCode < 400 {
				sum.SuccessRequests++
			} else {
				sum.FailedRequests++

			}
			sum.StatusCodes[res.StatusCode]++
		}

		if res.Duration > 0 {
			durations = append(durations, res.Duration);
			totalDurationSum += res.Duration;
		}
	}

	if totalDuration.Seconds() > 0 {
		sum.RPS = float64(sum.TotalRequests) / totalDuration.Seconds();
	}

	if len(durations) > 0 {
		return sum;
	}

	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	sum.MinDuration = durations[0];
	sum.MaxDuration = durations[len(durations) - 1];
	sum.AvgDuration = totalDurationSum / time.Duration(len(durations));

	sum.P50 = calculatePercentile(durations, 0.50);
	sum.P95 = calculatePercentile(durations, 0.95);
	sum.P99 = calculatePercentile(durations, 0.99);

	return sum
}
