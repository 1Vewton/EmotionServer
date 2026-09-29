package emotionapi

// NewContextQuery creates new context
type NewContextQuery struct {
	APIKey string `json:"api_key" form:"api_key"`
}

// WithContextResponse returns response with context key
type WithContextResponse struct {
	ContextKey string `json:"context_key" form:"context_key"`
}
