package cmd

import (
	"fmt"
	"os"

	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [URL of your Ebooks]",
	Short: "downloads your ebooks from a url (order link)",
	Long:  "downloads your ebooks from a url. Use this, if your order is public and not claimed",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No url provided")
			os.Exit(-1)
		}
		fmt.Printf("Load ebooks from %s\n", args[0])
		doc, err := htmlquery.LoadURL(args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		query := "//div[contains(@class, 'download-buttons')]//div[contains(@class,'js-start-download')]//a"

		for _, book := range htmlquery.Find(doc, query) {
			for _, link := range book.Attr {
				if link.Key == "href" {
					downloadFile(link.Val, args[1])
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVar(&outputDirectory, "outputDirectory", ".", "Output location for downloaded ebooks")

}
