package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// testRequests tests the request
func testRequests(
	engine *gin.Engine,
	method string,
	url string,
	data any,
) (*response.Response, error) {
	w := httptest.NewRecorder()
	var reqBuffer io.Reader = nil
	if data == nil {
		reqBuffer = nil
	} else {
		bytesData, err := json.Marshal(
			data,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"%s, original : %s",
				err.Error(),
				string(bytesData),
			)
		}
		reqBuffer = bytes.NewBuffer(bytesData)
	}
	req, err := http.NewRequest(
		method,
		url,
		reqBuffer,
	)
	if err != nil {
		return nil, err
	}
	engine.ServeHTTP(
		w,
		req,
	)
	resultBytes := w.Body.Bytes()
	var result *response.Response
	err = json.Unmarshal(
		resultBytes,
		&result,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"%s, original : %s",
			err.Error(),
			string(resultBytes),
		)
	}
	return result, nil
}

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
			resp, err := testRequests(
				router,
				"GET",
				"/v1/utils/health",
				nil,
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
					resp, err := testRequests(
						router,
						"POST",
						"/v1/data/addAgentProfile",
						newAgentData,
					)
					if err != nil {
						t.Fatal(err)
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
							"the resp data (%T) with key cannot be decoded , raw data: %v",
							resp.Data,
							resp,
						)
					}
					rawKey, exists := respWithKey["AgentKey"]
					if !exists {
						t.Error(
							"cannot find the agent key returned",
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
					newQueryData := dataapi.GetAgentProfileQuery{
						AgentKey: key,
					}
					resp, err := testRequests(
						router,
						"POST",
						"/v1/data/getAgentProfile",
						newQueryData,
					)
					if err != nil {
						t.Errorf(
							"data: %v",
							resp,
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
			resp, err := testRequests(
				router,
				"POST",
				"/v1/data/addAgentProfile",
				newAgentData,
			)
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
			agentKey, ok := processedData["AgentKey"].(string)
			if !ok {
				t.Fatalf(
					"fetched data %v of agent key cannot be converted to string",
					processedData["AgentKey"],
				)
			}
			parameter := &emotionapi.NewContextQuery{
				AgentKey: agentKey,
			}
			respCtx, err := testRequests(
				router,
				"POST",
				"/v1/emotion/registerEmotion",
				parameter,
			)
			if err != nil {
				t.Errorf(
					"%v",
					respCtx,
				)
				t.Fatal(err)
			}
			if !respCtx.Success {
				t.Fatalf(
					"the context generation failed due to %s",
					*respCtx.Error,
				)
			}
		},
	)
	t.Run(
		"TestGetEmotions",
		func(t *testing.T) {
			resp, err := testRequests(
				router,
				"POST",
				"/v1/data/addAgentProfile",
				newAgentData,
			)
			if err != nil {
				t.Fatal(err)
			}
			if !resp.Success {
				if resp.Error == nil {
					t.Fatal(
						"unexpected error",
					)
				}
				t.Fatalf(
					"the agent creating failed due to %s",
					*resp.Error,
				)
			}
			newMap, canConvert := resp.Data.(map[string]any)
			if !canConvert {
				t.Fatal(
					"the data cannot be converted to map",
				)
			}
			data, exists := newMap["AgentKey"]
			if !exists {
				t.Fatal(
					"the AgentKey key does not exists",
				)
			}
			agentKey, isString := data.(string)
			if !isString {
				t.Fatal(
					"the AgentKey cannot simply converted to string",
				)
			}
			// Add multiple contexts
			addContextParameter := &emotionapi.NewContextQuery{
				AgentKey: agentKey,
			}
			resp, err = testRequests(
				router,
				"POST",
				"/v1/emotion/registerEmotion",
				addContextParameter,
			)
			if err != nil {
				t.Fatal(err)
			}
			if !resp.Success {
				if resp.Error == nil {
					t.Fatal(
						"unexpected error",
					)
				}
				t.Fatal(
					*resp.Error,
				)
			}
			resp, err = testRequests(
				router,
				"POST",
				"/v1/emotion/registerEmotion",
				addContextParameter,
			)
			if err != nil {
				t.Fatal(err)
			}
			if !resp.Success {
				if resp.Error == nil {
					t.Fatal(
						"unexpected error",
					)
				}
				t.Fatal(
					*resp.Error,
				)
			}
			// Get contexts
			getContextsParameter := emotionapi.GetContextsOfAProfileQuery{
				AgentKey: agentKey,
			}
			resp, err = testRequests(
				router,
				"POST",
				"/v1/emotion/getContextsOfAProfile",
				getContextsParameter,
			)
			if err != nil {
				t.Fatal(err)
			}
			if !resp.Success {
				if resp.Error != nil {
					t.Fatal(*resp.Error)
				}
				t.Fatal(
					"unexpected error",
				)
			}
			mapData, isMap := resp.Data.(map[string]any)
			if !isMap {
				t.Fatal(
					"the fetched data is not a map",
				)
			}
			contextsRaw, exists := mapData["Contexts"]
			if !exists {
				t.Fatal(
					"the returned data do not contain context field",
				)
			}
			_, isStringSlice := contextsRaw.([]any)
			if !isStringSlice {
				t.Fatalf(
					"the contexts is not a string slice: %v, all data: %v",
					contextsRaw,
					resp,
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
