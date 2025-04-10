package ai

import (
	"app/internal/config"
	"app/internal/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	cfg *config.Config
}

func NewAIHandler(cfg *config.Config) *AIHandler {
	return &AIHandler{
		cfg: cfg,
	}
}

// RegisterRoutes registers all AI-related routes
func (h *AIHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/ai", h.handleAIEndpoint)
}

// handleAIEndpoint handles the AI endpoint request
// @Summary: Handles AI requests and forwards them to the Gemini API
// @Description: Receives a question, forwards it to the Gemini API, and returns the answer
// @Tags: AI
// @Accept: json
// @Produce: json
// @Param request body domain.AIRequest true "AI Request"
// @Success 200 {object} domain.AIResponse
// @Failure 400 {object} fiber.Map "Invalid request body"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /ai [post]
func (h *AIHandler) handleAIEndpoint(c *fiber.Ctx) error {
	// Parse request body
	var req domain.AIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate request
	if req.Question == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Question is required",
		})
	}

	// Prepare Gemini API request
	geminiReq := domain.GeminiRequest{
		Contents: []domain.GeminiContent{
			{
				Parts: []domain.GeminiPart{
					{
						Text: req.Question,
					},
				},
			},
		},
	}

	// Convert request to JSON
	reqBody, err := json.Marshal(geminiReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to prepare request",
		})
	}

	// Create HTTP request to Gemini API
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		h.cfg.AIProvider.Gemini.Model, h.cfg.AIProvider.Gemini.ApiKey)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request",
		})
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Send request to Gemini API
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to call Gemini API: %v", err),
		})
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response",
		})
	}

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Gemini API error (status %d): %s", resp.StatusCode, string(respBody)),
		})
	}

	// Parse Gemini API response
	var geminiResp domain.GeminiResponse
	if err := json.Unmarshal(respBody, &geminiResp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to parse Gemini API response: %v", err),
		})
	}

	// Extract answer from response
	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No response from Gemini API",
		})
	}

	answer := geminiResp.Candidates[0].Content.Parts[0].Text

	// Return response
	return c.JSON(domain.AIResponse{
		Answer: answer,
	})
}
