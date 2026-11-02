package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1Vewton/EmotionServer/api/dataapi"
	"github.com/1Vewton/EmotionServer/api/emotionapi"
	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/1Vewton/EmotionServer/internal/manager"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/mitchellh/mapstructure"
)

// TestRouter tests functions related to router
func TestRouter(t *testing.T) {
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
	newAgentData := &dataapi.NewAgentProfileQuery{
		Openness:          0.0,
		Conscientiousness: 0.0,
		Extraversion:      0.0,
		Agreeableness:     0.0,
		Neuroticism:       0.0,
		IsContextInfinite: true,
	}
	redisclient := database.NewRedisConfig(
		"localhost:6379",
		"",
	).WithDialTimeout(
		5,
	).WithReadTimeout(
		5,
	).WithMaxRetries(
		3,
	).WithMaxRetryBackoff(
		5,
	).WithMinRetryBackoff(
		3,
	).ToClient()
	manager.MainContextManager = manager.NewContextManager(
		redisclient,
		manager.MainAgentProfileManager,
	)
	// TestCheckHealth tests the check health endpoint of the api
	t.Run(
		"TestCheckHealth",
		func(t *testing.T) {
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
		},
	)
	// TestAgentProfileCRUD tests the CRUD of agent profile
	t.Run(
		"TestAgentProfileCRUD",
		func(t *testing.T) {
			var key string = ""
			// Creation test
			t.Run(
				"DataCreationTest",
				func(t *testing.T) {
					w := httptest.NewRecorder()
					encodedData, err := json.Marshal(newAgentData)
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
					rawKey, exists := respWithKey["APIKey"]
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
		},
	)
	t.Run(
		"TestEmotionCRUD",
		func(t *testing.T) {
			// Create the agent profile and creates the context
			w := httptest.NewRecorder()
			encodedData1, err := json.Marshal(
				newAgentData,
			)
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(
				"POST",
				"/v1/data/addAgentProfile",
				bytes.NewBuffer(encodedData1),
			)
			if err != nil {
				t.Fatal(err)
			}
			router.ServeHTTP(
				w,
				req,
			)
			result := w.Body.Bytes()
			var resp response.Response
			err = json.Unmarshal(result, &resp)
			if err != nil {
				t.Fatal(err)
			}
			if !resp.Success {
				t.Fatalf(
					"the request is not successful due to %s",
					*resp.Error,
				)
			}
			processedData, ok := resp.Data.(map[string]any)
			if !ok {
				t.Fatalf(
					"expected to be allowed to decode to new agent profile, resp data: %v",
					resp.Data,
				)
			}
			apiKey, ok := processedData["APIKey"].(string)
			if !ok {
				t.Fatalf(
					"fetched data %v of api key cannot be converted to string",
					processedData["APIKey"],
				)
			}
			parameter := &emotionapi.NewContextQuery{
				APIKey: apiKey,
			}
			decodedDataParameter, err := json.Marshal(parameter)
			if err != nil {
				t.Fatal(err)
			}
			w2 := httptest.NewRecorder()
			req, err = http.NewRequest(
				"POST",
				"/v1/emotion/registerEmotion",
				bytes.NewBuffer(decodedDataParameter),
			)
			if err != nil {
				t.Fatal(err)
			}
			router.ServeHTTP(
				w2,
				req,
			)
			ctxResult := w2.Body.Bytes()
			var respCtx response.Response
			err = json.Unmarshal(
				ctxResult,
				&respCtx,
			)
			if err != nil {
				t.Error(
					string(ctxResult),
				)
				t.Fatal(err)
			}
			if !respCtx.Success {
				t.Fatalf(
					"the context generation is insuccess due to %s",
					*respCtx.Error,
				)
			}
		},
	)
	err = database.Close(db)
	if err != nil {
		t.Fatal(err)
	}
	err = database.CloseRedis(redisclient)
	if err != nil {
		t.Fatal(err)
	}
}
