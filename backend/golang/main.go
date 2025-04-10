package main

import (
	"app/cmd"
	_ "app/docs"
	"log/slog"
	"os"
)

// @title Coding Test API
// @version 1.0.0
// @description This is a sample API server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		slog.Error("Failed to execute command", "error", err)
		// Exit with a non-zero status code
		os.Exit(1)
	}
}
