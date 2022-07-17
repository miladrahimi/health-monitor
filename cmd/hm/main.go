package main

import (
    "fmt"
    "github.com/miladrahimi/health-monitor/internal/app"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "HM",
    Short: "Health Monitor",
    Long:  `Health Monitor calls health-check endpoints and draws a timing chart.`,
    Run: func(cmd *cobra.Command, args []string) {
        app.Serve()
    },
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
        os.Exit(1)
    }
}
