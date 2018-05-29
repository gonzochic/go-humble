package cmd

import (
	"fmt"
	"os"

	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
)

// getCmd represents the get command

var getLocal = &cobra.Command{
	Use:   "local [PATH to HTML of Ebooks]",
	Short: "downloads your ebooks from local website copy",
	Long:  "downloads your ebooks from your local disk. Use this, if your order is claimed by a user",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Load ebooks from %s\n", args[0])
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		doc, err := htmlquery.Parse(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		query := "//div[contains(@class, 'download-buttons')]//div[contains(@class,'js-start-download')]//a"

		for _, book := range htmlquery.Find(doc, query) {
			for _, link := range book.Attr {
				if link.Key == "href" {
					downloadFile(link.Val, outputDirectory)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getLocal)
	getLocal.Flags().StringVar(&outputDirectory, "outputDirectory", ".", "Output location for downloaded ebooks")
}
