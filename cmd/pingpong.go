package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	projectVersion string
	projectCommit  string
	displayVersion bool
	logLevel       string
)

func init() {
	PingPongCommand.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Configures the log level to use (debug, info, warn, error, fatal, panic)")
	PingPongCommand.Flags().BoolVarP(&displayVersion, "version", "v", false, "Display the current version of pingpong")

	PingPongCommand.AddCommand(versionCommand)
	PingPongCommand.AddCommand(serverCommand)
	PingPongCommand.AddCommand(clientCommand)
}

var PingPongCommand = &cobra.Command{
	Use:   "pingpong",
	Short: "Pingpong is a simple request/response test tool",
	Long:  `Pingpong can run in server and client mode.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if _, err := logrus.ParseLevel(logLevel); err != nil {
			logrus.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if displayVersion {
			printProjectVersion()
		} else {
			cmd.Help()
		}
	},
}
