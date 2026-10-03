package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1Vewton/EmotionServer/api/dataapi"
	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/1Vewton/EmotionServer/internal/manager"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/mitchellh/mapstructure"
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
	var resp response.Response
	err = json.Unmarshal(
		result,
		&resp,
	)
	if err != nil {
		t.Error(err)
	}
	if !resp.Success {
		if resp.Error == nil {
			t.Error("no error info displayed")
		} else {
			t.Error(*resp.Error)
		}
	}
}

// TestAgentProfileCRUD tests the CRUD of agent profile
func TestAgentProfileCRUD(t *testing.T) {
	t.Parallel()
	// Initialization
	db, err := database.Connect(
		":memory:",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		panic(err)
	}
	manager.MainAgentProfileManager = manager.NewAgentProfileManager(
		db,
	)
	router := SetUpRouter()
	var key string = ""
	newData := dataapi.NewAgentProfileQuery{
		Openness:          0.0,
		Conscientiousness: 0.0,
		Extraversion:      0.0,
		Agreeableness:     0.0,
		Neuroticism:       0.0,
		IsContextInfinite: true,
	}
	// Creation test
	t.Run(
		"DataCreationTest",
		func(t *testing.T) {
			w := httptest.NewRecorder()
			encodedData, err := json.Marshal(newData)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(
				"POST",
				"/v1/data/addAgentProfile",
				bytes.NewBuffer(encodedData),
			)
			if err != nil {
				t.Error(err)
			}
			router.ServeHTTP(w, req)
			result := w.Body.Bytes()
			var resp response.Response
			err = json.Unmarshal(
				result,
				&resp,
			)
			if err != nil {
				t.Errorf(
					"data: %s",
					string(result),
				)
				t.Error(err)
			}
			if !resp.Success {
				if resp.Error == nil {
					t.Error(
						"unable to show error info",
					)
				} else {
					t.Error(
						*resp.Error,
					)
				}
			}
			var respWithKey map[string]any
			respWithKey, ok := resp.Data.(map[string]any)
			if !ok {
				t.Errorf(
					"the resp data (%T) with key cannot be decoded , raw data: %s",
					resp.Data,
					string(result),
				)
			}
			rawKey, exists := respWithKey["api_key"]
			if !exists {
				t.Error(
					"cannot find the api key returned",
				)
			}
			key, ok = rawKey.(string)
			if !ok {
				t.Error(
					"the key returned is not a string",
				)
			}
			t.Log(key)
		},
	)
	// Get Test
	t.Run(
		"GetAgentProfileTest",
		func(t *testing.T) {
			w := httptest.NewRecorder()
			newQueryData := dataapi.GetAgentProfileQuery{
				APIKey: key,
			}
			encodedData, err := json.Marshal(
				newQueryData,
			)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(
				"POST",
				"/v1/data/getAgentProfile",
				bytes.NewBuffer(encodedData),
			)
			router.ServeHTTP(
				w,
				req,
			)
			result := w.Body.Bytes()
			var resp response.Response
			err = json.Unmarshal(
				result,
				&resp,
			)
			if err != nil {
				t.Errorf(
					"data: %s",
					string(result),
				)
				t.Error(err)
			}
			if !resp.Success {
				if resp.Error == nil {
					t.Error(
						"unable to show error info",
					)
				} else {
					t.Error(
						*resp.Error,
					)
				}
			}
			var resultProfile *profile.AgentProfile
			err = mapstructure.Decode(
				resp.Data,
				&resultProfile,
			)
			if err != nil {
				t.Error(err)
			}
		},
	)
	err = database.Close(db)
	if err != nil {
		t.Error(err)
	}
}
