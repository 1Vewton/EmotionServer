package ocean

// Personality stores the OCEAN personality
type Personality struct {
	Openness          float64
	Conscientiousness float64
	Extraversion      float64
	Agreeableness     float64
	Neuroticism       float64
}

// NewPersonality creates new personality
func NewPersonality(
	openness float64,
	conscientiousness float64,
	extraversion float64,
	agreeableness float64,
	neuroticism float64,
) *Personality {
	return &Personality{
		Openness:          openness,
		Conscientiousness: conscientiousness,
		Extraversion:      extraversion,
		Agreeableness:     agreeableness,
		Neuroticism:       neuroticism,
	}
}
