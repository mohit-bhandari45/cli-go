package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mohit-bhandari45/gurl/internal/runner"
	"github.com/mohit-bhandari45/gurl/internal/stats"
	"github.com/spf13/cobra"
)

// methods to hold flag values
var (
	method      string
	headers  	[]string
	data		string
	totalReqs   int
	concurrency int
	output 		string
)

var rootCmd = &cobra.Command{
	Use:   "gurl [url]",
	Short: "gurl is a concurrent HTTP client and API load testing tool",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetURL := args[0]

		cfg := runner.RequestConfig {
			URL: targetURL,
			Method: method,
			Headers: parseHeaders(headers),
			Body: []byte(data),
			Timeout: 10 * time.Second,
			TotalReqs: totalReqs,
			Concurrency: concurrency,
		}

		start := time.Now();
		results := runner.Run(cfg)
		totalDuration := time.Since(start);
		summary := stats.Calculate(results, totalDuration);

		fmt.Println(totalDuration, summary);
		fmt.Println(totalDuration.Seconds(), summary);

		return nil
	},
}

func parseHeaders(headers []string) http.Header {
	h := make(http.Header);
	
	for _, header := range headers {
		parts := strings.SplitN(header, ":", 2);
		if len(parts) == 2 {
			h.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]));
		}
	}

	return h;
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// StringVarP binds a string flag supporting both long (--method) and short POSIX (-X) shorthand formats into a target variable pointer.
	
	rootCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP method (GET, POST, etc.)");
	rootCmd.Flags().StringSliceVarP(&headers, "header", "H", []string{}, "Custom HTTP headers (e.g. -H 'Content-Type: application/json')");
	rootCmd.Flags().StringVarP(&data, "data", "d", "", "HTTP request body payload");
	rootCmd.Flags().IntVarP(&totalReqs, "requests", "n", 100, "Total number of HTTP requests to send");
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Number of concurrent worker goroutines");
	rootCmd.Flags().StringVarP(&output, "output", "o", "text", "Output format (text or json)");
}