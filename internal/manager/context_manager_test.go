package manager

import (
	"testing"
)

// inList checks if certain string is in a string list
func inList(
	t *testing.T,
	content string,
	singleString string,
	stringList []string,
) {
	for _, selectedString := range stringList {
		if selectedString == singleString {
			return
		}
	}
	t.Fatalf(
		"%s does not exists in the list of %s",
		singleString,
		content,
	)
}

// notInList checks if certain string is not in a string list
func notInList(
	t *testing.T,
	content string,
	singleString string,
	stringList []string,
) {
	for _, selectedString := range stringList {
		if selectedString == singleString {
			t.Fatalf(
				"%s exists in the list of %s",
				singleString,
				content,
			)
		}
	}
}
