package cmd

import (
	"github.com/spf13/cobra"
)

var (
	projectVersion string
	projectCommit  string
	displayVersion bool
)

func init() {
	PingPongCommand.Flags().BoolVarP(&displayVersion, "version", "v", false, "Display the current version of pingpong")

	PingPongCommand.AddCommand(versionCommand)
	PingPongCommand.AddCommand(serverCommand)
	PingPongCommand.AddCommand(clientCommand)
}

var PingPongCommand = &cobra.Command{
	Use:   "pingpong",
	Short: "Pingpong is a simple request/response test tool",
	Long:  `Pingpong can run in server and client mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		if displayVersion {
			printProjectVersion()
		} else {
			cmd.Help()
		}
	},
}
