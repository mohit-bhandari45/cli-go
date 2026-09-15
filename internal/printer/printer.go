package printer

import (
	"encoding/json"
	"fmt"
	
	"github.com/fatih/color"
	"github.com/mohit-bhandari45/gurl/internal/stats"
)

func PrintJSON(sum stats.Summary) error {
	jsonData, err := json.MarshalIndent(sum, "", "   ");
	if err != nil {
		return fmt.Errorf("failed to marshal summary to JSON: %w", err)
	}

	fmt.Println(string(jsonData))
	return nil
}

func PrintText(sum stats.Summary, noColor bool) {
	if noColor {
		color.NoColor = true
	}

	cyan := color.New(color.FgCyan, color.Bold).SprintfFunc()
	green := color.New(color.FgGreen, color.Bold).SprintfFunc()
	red := color.New(color.FgRed, color.Bold).SprintfFunc()
	yellow := color.New(color.FgYellow, color.Bold).SprintfFunc()

	fmt.Println()
	fmt.Println(cyan("=== gurl Load Test Results ==="))
	fmt.Printf("Total Requests:     %s\n", yellow("%d", sum.TotalRequests))
	fmt.Printf("Successful:         %s\n", green("%d", sum.SuccessRequests))
	fmt.Printf("Failed:             %s\n", red("%d", sum.FailedRequests))
	fmt.Printf("Total Duration:     %v\n", sum.TotalDuration)
	fmt.Printf("Requests / Sec:     %s\n", green("%.2f", sum.RPS))
	
	fmt.Println()
	fmt.Println(cyan("--- Latency Distribution ---"))
	fmt.Printf("  Min:              %v\n", sum.MinDuration)
	fmt.Printf("  Mean (Average):   %v\n", sum.AvgDuration)
	fmt.Printf("  p50 (Median):     %v\n", sum.P50)
	fmt.Printf("  p95:              %v\n", sum.P95)
	fmt.Printf("  p99 (Tail):       %v\n", sum.P99)
	fmt.Printf("  Max:              %v\n", sum.MaxDuration)

	if len(sum.StatusCodes) > 0 {
		fmt.Println()
		fmt.Println(cyan("--- Status Code Breakdown ---"))
		for code, count := range sum.StatusCodes {
			if code >= 200 && code < 400 {
				fmt.Printf("  %s: %d\n", green("[%d]", code), count)
			} else {
				fmt.Printf("  %s: %d\n", red("[%d]", code), count)
			}
		}
	}

	if len(sum.Errors) > 0 {
		fmt.Println()
		fmt.Println(red("--- Error Breakdown ---"))
		for errStr, count := range sum.Errors {
			fmt.Printf("  %s: %d\n", errStr, count)
		}
	}
	fmt.Println()
}