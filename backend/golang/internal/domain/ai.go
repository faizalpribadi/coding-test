package domain

// Request structure for AI endpoint
type AIRequest struct {
	Question string `json:"question"`
}

// Response structure for AI endpoint
type AIResponse struct {
	Answer string `json:"answer"`
}

// Gemini API request structure
type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

// Gemini content structure
type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

// Gemini part structure
type GeminiPart struct {
	Text string `json:"text"`
}

// Gemini API response structure
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}
