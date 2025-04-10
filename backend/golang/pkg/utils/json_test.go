package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSONFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "json-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Test cases
	testCases := []struct {
		name          string
		jsonContent   string
		expectedData  map[string]interface{}
		expectedError bool
	}{
		{
			name: "Valid JSON file",
			jsonContent: `{
				"name": "John Doe",
				"age": 30,
				"email": "john@example.com"
			}`,
			expectedData: map[string]interface{}{
				"name":  "John Doe",
				"age":   float64(30),
				"email": "john@example.com",
			},
			expectedError: false,
		},
		{
			name:          "Invalid JSON file",
			jsonContent:   `{"name": "John Doe", "age": 30, "email":}`,
			expectedData:  nil,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempFile := filepath.Join(tempDir, "test.json")
			err := os.WriteFile(tempFile, []byte(tc.jsonContent), 0644)
			assert.NoError(t, err)

			var result map[string]interface{}
			err = LoadJSONFile(tempFile, &result)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedData, result)
			}
		})
	}
}
