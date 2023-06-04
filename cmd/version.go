/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of terraform-imgs",
	Long:  `Print the version number of terraform-imgs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
