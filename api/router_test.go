package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1Vewton/EmotionServer/api/response"
)

// TestCheckHealth tests the check health endpoint of the api
func TestCheckHealth(
	t *testing.T,
) {
	t.Parallel()
	// Start server
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(
		"GET",
		"/v1/utils/health",
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(w, req)
	result := w.Body.Bytes()
	var response response.Response
	err = json.Unmarshal(
		result,
		&response,
	)
	if err != nil {
		t.Error(err)
	}
	if !response.Success {
		if response.Error == nil {
			t.Error("no error info displayed")
		} else {
			t.Error(*response.Error)
		}
	}
}
