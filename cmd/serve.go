package cmd

import (
	"fmt"
	"lms-backend/cmd/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the LMS web server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server...")
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
