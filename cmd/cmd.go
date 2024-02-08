package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emuk",
	Short: "Emuk is a CPU and memory simulator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
