package sales

import (
	"app/pkg/utils"
	"path/filepath"

	"app/internal/config"
	"app/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type SalesHandler struct {
	cfg *config.Config
}

func NewSalesHandler(cfg *config.Config) *SalesHandler {
	return &SalesHandler{
		cfg: cfg,
	}
}

// RegisterRoutes registers all sales-related routes
func (h *SalesHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/data", h.handleGetData)
	router.Get("/sales-reps", h.handleGetSalesReps)
	router.Get("/sales-reps/filter", h.handleFilterSalesReps)
}

// handleGetData returns the full dummy data
// @Summary: Returns the full dummy data
// @Description: Returns the full dummy data from the JSON file
// @Tags: Sales
// @Accept: json
// @Produce: json
// @Success 200 {object} domain.DummyData
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /sales/data [get]
func (h *SalesHandler) handleGetData(c *fiber.Ctx) error {
	var data any

	// Get absolute path to the JSON file
	filePath, err := filepath.Abs("../../dummyData.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to resolve file path",
		})
	}

	// Load the JSON data
	if err := utils.LoadJSONFile(filePath, &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load data",
		})
	}

	return c.JSON(data)
}

// handleGetSalesReps returns all sales representatives
// @Summary: Returns all sales representatives
// @Description: Returns all sales representatives from the JSON file
// @Tags: Sales
// @Accept: json
// @Produce: json
// @Success 200 {object} []domain.SalesRep
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /sales/sales-reps [get]

func (h *SalesHandler) handleGetSalesReps(c *fiber.Ctx) error {
	var data domain.DummyData

	// Get absolute path to the JSON file
	filePath, err := filepath.Abs("../../dummyData.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to resolve file path",
		})
	}

	// Load the JSON data
	if err := utils.LoadJSONFile(filePath, &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load data",
		})
	}

	return c.JSON(data.SalesReps)
}

// handleFilterSalesReps filters sales representatives by region and/or deal status
// @Summary: Filters sales representatives by region and/or deal status
// @Description: Filters sales representatives by region and/or deal status from the JSON file
// @Tags: Sales
// @Accept: json
// @Produce: json
// @Param region query string false "Region to filter by"
// @Param deal_status query string false "Deal status to filter by"
// @Success 200 {object} []domain.SalesRep
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /sales/sales-reps/filter [get]
func (h *SalesHandler) handleFilterSalesReps(c *fiber.Ctx) error {
	// Get query parameters
	region := c.Query("region")
	dealStatus := c.Query("deal_status")

	// If no filters provided, return all sales reps
	if region == "" && dealStatus == "" {
		return h.handleGetSalesReps(c)
	}

	var data domain.DummyData

	// Get absolute path to the JSON file
	filePath, err := filepath.Abs("../../dummyData.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to resolve file path",
		})
	}

	// Load the JSON data
	if err := utils.LoadJSONFile(filePath, &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load data",
		})
	}

	// Filter sales reps
	var filteredReps []domain.SalesRep

	// Apply filters
	for _, rep := range data.SalesReps {
		// Filter by region if provided
		if region != "" && rep.Region != region {
			continue
		}

		// Filter by deal status if provided
		if dealStatus != "" {
			// Check if rep has any deals with the specified status
			hasDealWithStatus := false
			var filteredDeals []domain.Deal

			for _, deal := range rep.Deals {
				if deal.Status == dealStatus {
					hasDealWithStatus = true
					filteredDeals = append(filteredDeals, deal)
				}
			}

			// Skip rep if they don't have any deals with the specified status
			if !hasDealWithStatus {
				continue
			}

			// Add filtered deals to the rep
			rep.FilteredDeals = filteredDeals
		}

		// Add rep to filtered list
		filteredReps = append(filteredReps, rep)
	}

	return c.JSON(filteredReps)
}
