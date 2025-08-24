package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "loglet",
	Short: "Fast, minimalist static blog generator",
	Long:  "A fast, minimalist static blog generator that turns plain text into clean, timeless HTML in seconds",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
