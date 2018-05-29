package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var outputDirectory string

var rootCmd = &cobra.Command{
	Use:   "humble",
	Short: "Humble fetches your humble bundle ebooks",
	Long:  `Humble fetches your humble bundle ebooks`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
