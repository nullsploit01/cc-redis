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
	Long: `ccredis-server initiates a custom Redis server tailored for enhanced performance and reliability. The server configuration allows for customization, including specifying a custom port to suit various deployment environments. 

		Examples of using this application include:

		- Launching the server with default settings (host: localhost, port: 6379):
		$ ccredis-server

		- Specifying a custom port for the server:
		$ ccredis-server --port 6380

		This tool provides a user-friendly interface for configuring and managing your Redis server, enabling easy deployment and flexibility.`,

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
