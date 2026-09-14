package runner

import "time"

// result to store the results
type Result struct {
	StatusCode 		int					`json:"status_code"`
	Duration		time.Duration		`json:"duration_ns"`
	BytesRead		int64				`json:"bytes_read"`
	Error 			string				`json:"error,omitempty"`
}