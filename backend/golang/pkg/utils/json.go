package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadJSONFile loads a JSON file and convert it into the provided interface
func LoadJSONFile(filePath string, target any) error {
	// Get the absolute path of the file
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Read the file
	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file %s: %w", absPath, err)
	}

	// Unmarshal the JSON data
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return nil
}
