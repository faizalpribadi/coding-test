package cmd

import (
	"app/internal/api/ai"
	"app/internal/api/sales"
	"app/internal/config"
	"fmt"

	_ "app/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  `Start the server http api`,
	RunE:  startServer,
}

func startServer(cmd *cobra.Command, args []string) error {
	// Load the configuration file
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return err
	}

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "Coding Test API",
	})

	// Add middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Set up Swagger documentation
	app.Get("/docs/*", swagger.HandlerDefault) // default

	// Setup API routes
	api := app.Group("/api")

	// Register sales routes
	salesHandler := sales.NewSalesHandler(cfg)
	salesHandler.RegisterRoutes(api)

	// Register AI routes
	aiHandler := ai.NewAIHandler(cfg)
	aiHandler.RegisterRoutes(api)

	// Start the server
	return app.Listen(fmt.Sprintf(":%d", cfg.App.Port))
}
