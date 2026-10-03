package dataapi

// NewAgentProfileQuery defines the data structure for creating new agent profile
type NewAgentProfileQuery struct {
	Openness              float64 `json:"openness" form:"openness"`
	Conscientiousness     float64 `json:"conscientiousness" form:"conscientiousness"`
	Extraversion          float64 `json:"extraversion" form:"extraversion"`
	Agreeableness         float64 `json:"agreeableness" form:"agreeableness"`
	Neuroticism           float64 `json:"neuroticism" form:"neuroticism"`
	IsContextInfinite     bool    `json:"is_context_infinite" form:"is_context_infinite"`
	ContextLifeTimeInDays *int    `json:"context_life_time_in_days" form:"context_life_time_in_days"`
}

// NewAgentProfileResponse defines the response for creating new agent profile response
type NewAgentProfileResponse struct {
	APIKey string `json:"api_key"`
}

// GetAgentProfileQuery defines the query for searching the agent profile
type GetAgentProfileQuery struct {
	APIKey string `json:"api_key"`
}
