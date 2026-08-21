package mathematics

import (
	"math"
)

// RoundDigits round the data with certain digits
func RoundDigits(
	data float64,
	digit float64,
) float64 {
	return math.Round(data/digit) * digit
}
