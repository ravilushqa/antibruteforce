package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	subnet  string
	address string
	login   string
	ip      string
)

var rootCmd = &cobra.Command{
	Use:   "antibruteforce",
	Short: "antibruteforce service",
	Long:  `antibruteforce service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Use antibruteforce [command]\nRun 'antibruteforce --help' for usage.\n")
	},
}

// Execute is starting application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("root execute error: %v", err)
	}
}
