/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"github.com/spf13/cobra"
)

// SouceCmd represents the souce command
var SourceCmd = &cobra.Command{
	Use:   "source",
	Short: "The source contains source processing tools.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// souceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// souceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
