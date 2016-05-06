package cmd

import "github.com/spf13/cobra"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pingpong",
	Long:  `Print the version number of pingpong.`,
	Run: func(cmd *cobra.Command, args []string) {
		printProjectVersion()
	},
}
