package mathematics

import (
	"testing"
)

// Test the round digit process
func TestRoundDigit(
	t *testing.T,
) {
	const testVal float64 = 0.114514
	testVal1 := RoundDigits(testVal, 0.01)
	if testVal1 != 0.11 {
		t.Errorf(
			"Expected %f, got %f",
			testVal1,
			0.11,
		)
	}
	testVal2 := RoundDigits(testVal, 0.001)
	if testVal2 != 0.115 {
		t.Errorf(
			"Expected %f, got %f",
			testVal2,
			0.115,
		)
	}
}
