package cmd

import (
	log "github.com/Sirupsen/logrus"
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
		switch logLevel {
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warn":
			log.SetLevel(log.WarnLevel)
		case "error":
			log.SetLevel(log.ErrorLevel)
		case "fatal":
			log.SetLevel(log.FatalLevel)
		case "panic":
			log.SetLevel(log.PanicLevel)
		default:
			log.Fatalf("Cannot use unsupported log level %s.", logLevel)
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
