package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emuk [command]",
	Short: "Emuk is a CPU and memory simulator",
	Args:  cobra.ExactArgs(1),
	//Example: "aaa",
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
