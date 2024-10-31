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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

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
