package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dataServer",
	Short: "A brief description of your application",
	Long:  `Use dataServer run -h to see what this does`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Root cmd execute err %x\n", err)
	}
}

func init() {

}
