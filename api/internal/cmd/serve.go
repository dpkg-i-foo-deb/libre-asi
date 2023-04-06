package cmd

import (
	"fmt"
	"libre-asi-api/cmd/app"
	"libre-asi-api/cmd/routes"
	"libre-asi-api/util"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Libre-ASI API server",
	Long: `Start the libre-asi-api server to handle requests from clients and provide
access to the Addiction Severity Index (ASI) form data. The server listens
for incoming requests on the specified port.`,
	Args: validateServeArgs,
	Run:  runServe,
}

func init() {
	serveCmd.Flags().IntP("port", "p", 3000, "Server Port")
	rootCmd.AddCommand(serveCmd)
}

func validateServeArgs(cmd *cobra.Command, args []string) error {

	// Ensure that the port argument is a valid integer value
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}
	if port < 0 || port > 65535 {
		return fmt.Errorf("invalid port value: %d (must be between 0 and 65535)", port)
	}
	return nil

}

func runServe(cmd *cobra.Command, args []string) {
	port, err := cmd.Flags().GetInt("port")

	if err != nil {
		util.HandleErrorStop(err)
	}

	server := app.SetUp()

	routes.SetRoutes(server)

	app.Start(port)
}
