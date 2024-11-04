package cmd

import (
	"os"

	"github.com/nullsploit01/cc-redis/cli/internal"
	"github.com/spf13/cobra"
)

var host string
var port string

var rootCmd = &cobra.Command{
	Use:   "ccredis-cli",
	Short: "Command line interface for managing custom Redis server",
	Long: `ccredis-cli is a command line interface designed to manage and interact with a custom Redis server. It provides comprehensive tools for monitoring, configuring, and controlling the server, ensuring optimal performance and ease of use.
		Examples of using ccredic-cli include:

		- Connecting to the Redis server on the default host and port:
		$ ccredic-cli

		- Specifying a custom port and host for the server:
		$ ccredic-cli --port 6380 --host 192.168.1.10`,

	Run: func(cmd *cobra.Command, args []string) {
		c := internal.NewCLI()
		err := c.ConnectToServer(host, port)
		if err != nil {
			cmd.PrintErrln(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&port, "port", "p", "6379", "Redis server port")
	rootCmd.Flags().StringVarP(&host, "host", "l", "localhost", "Redis server host")
}
