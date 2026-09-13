package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// methods to hold flag values
var (
	method      string
	totalReqs   int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "gurl [url]",
	Short: "gurl is a concurrent HTTP client and API load testing tool",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetURL := args[0]
		fmt.Printf("Target URL: %s\n", targetURL)
		fmt.Printf("Method: %s\n", method)
		fmt.Printf("Total Requests: %d\n", totalReqs)
		fmt.Printf("Concurrency: %d\n", concurrency)
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// StringVarP binds a string flag supporting both long (--method) and short POSIX (-X) shorthand formats into a target variable pointer.
	
	rootCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP method (GET, POST, etc.)");
	rootCmd.Flags().IntVarP(&totalReqs, "requests", "n", 100, "Total number of HTTP requests to send");
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Number of concurrent worker goroutines");
}