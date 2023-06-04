package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version string

	outputFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terraform-imgs",
	Short: "A utility to generate MERMAID code from terraform code using OPENAI API.\n",
	Long:  `A utility to generate MERMAID code from terraform code using OPENAI API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "o", "", "output file path(e.g. README.md)")
}
