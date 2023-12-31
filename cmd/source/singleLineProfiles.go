/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	sourcePath string
)

// singleLineProfilesCmd represents the singleLineProfiles command
var singleLineProfilesCmd = &cobra.Command{
	Use:   "singleLineProfiles",
	Short: "Reshape the XML format of Profiles and Permissionsets",
	Long: `Puts multiple lines of the XML in single lines so changes can be better viewed in 
version control systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("singleLineProfiles called")
	},
}

func init() {
	singleLineProfilesCmd.Flags().StringVarP(&sourcePath, "source", "s", "", "Main project source path")

	if err := singleLineProfilesCmd.MarkFlagRequired("source"); err != nil {
		fmt.Println(err)
	}

	SourceCmd.AddCommand(singleLineProfilesCmd)

}
