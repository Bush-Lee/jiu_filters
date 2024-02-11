package cmd

import (
	logic "dataServer/cmd/logic"

	"github.com/spf13/cobra"
)

// const CSV_URL = "https://docs.google.com/spreadsheets/d/e/2PACX-1vSWlferwOudNFwTBRNNSbZgdOZrfovhCKH5Kb83gfNQt0VVkIW8gH8VvVqKBuPIXxmKk_mZBBdsw_OA/pub?output=csv"
const CSV_URL = "https://docs.google.com/spreadsheets/d/e/2PACX-1vRuS3ElvxXQnCqnddPgYxOIMiuthx8UmlTAVJMDRkfjXq4AZJuRvwRn2rGqG9q7E-b_5AnHGTMafM4j/pub?gid=0&single=true&output=csv"

var generateCmd = &cobra.Command{
	Use:   "run",
	Short: "Creates a local server for local prototyping",
	Long: `Creates a local server for local prototyping.
For example:
	dataServer run -p 8080 -s ./my/path -u "some_new_url"
	`,
	Run: runServer,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("port", "p", 8081, "Port to run the server on")
	generateCmd.Flags().StringP("url", "u", CSV_URL, "URL to get data from")
	generateCmd.Flags().StringP("folder_path", "f", "./static/", "Static folder path")
	generateCmd.Flags().BoolP("data", "d", false, "Get data only")
}

func runServer(cmd *cobra.Command, args []string) {
	port, err := cmd.Flags().GetInt("port")
	if err != nil {

	}

	url, err := cmd.Flags().GetString("url")
	if err != nil {

	}

	static_path, err := cmd.Flags().GetString("folder_path")
	if err != nil {

	}

	if static_path[len(static_path)-1] != '/' {
		static_path = static_path + "/"
	}

	logic.GenStaticFolder(url, static_path)
	if dataOnly, _ := cmd.Flags().GetBool("data"); !dataOnly {
		logic.StartServer(static_path, port)
	}
}
