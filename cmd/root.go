package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gurl [url]",
	Short: "gurl is a concurrent HTTP client and API load testing tool",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetUrl := args[0];
		fmt.Printf("Target URL received: %s\n", targetUrl);
		return nil;
	},
}

func Execute() error {
	return rootCmd.Execute();
}