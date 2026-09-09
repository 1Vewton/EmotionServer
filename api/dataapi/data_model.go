package dataapi

// NewAgentProfileQuery defines the data structure for creating new agent profile
type NewAgentProfileQuery struct {
	APIKey            string  `json:"api_key" form:"api_key"`
	Openness          float64 `json:"openness" form:"openness"`
	Conscientiousness float64 `json:"conscientiousness" form:"conscientiousness"`
	Extraversion      float64 `json:"extraversion" form:"extraversion"`
	Agreeableness     float64 `json:"agreeableness" form:"agreeableness"`
	Neuroticism       float64 `json:"neuroticism" form:"neuroticism"`
}
