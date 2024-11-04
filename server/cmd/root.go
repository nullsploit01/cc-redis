package cmd

import (
	"os"

	"github.com/nullsploit01/cc-redis/server/internal"
	"github.com/spf13/cobra"
)

var port string

var rootCmd = &cobra.Command{
	Use:   "ccredis-server",
	Short: "Launches a custom Redis server",
	Long: `This command initiates a custom Redis server tailored for enhanced performance and reliability. The server configuration allows for detailed customization including port adjustments to suit various deployment environments. 
		Examples of using this application include:

		- Launching the server with default settings.
		- Setting a custom port for the server operations.

		Cobra is a powerful CLI library for Go that simplifies creating command line applications. This tool leverages Cobra to provide a robust, easy-to-use interface for running and managing your custom Redis server.`,

	Run: func(cmd *cobra.Command, args []string) {
		s := internal.InitServer(port)
		err := s.StartServer()
		if err != nil {
			cmd.PrintErrln(err.Error())
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
	rootCmd.Flags().StringVarP(&port, "port", "p", "6379", "Port to run server on")
}
