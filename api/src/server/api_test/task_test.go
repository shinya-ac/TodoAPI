package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTask_GetTask(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		url          string
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"正常系": {
			url:          "/v1/tasks/",
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"tasks": []map[string]interface{}{
					{
						"Id":      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
						"Title":   "Todoのテストを行う1",
						"Content": "Todo機能のテストをGo言語で行う1",
						"Status":  "Pending",
					},
					{
						"Id":      "a74e2d5d-72ff-46c4-92ee-9b474c8c5588",
						"Title":   "Todoのテストを行う5",
						"Content": "Todo機能のテストをGo言語で行う5",
						"Status":  "InProgress",
					},
					{
						"Id":      "b6a38d18-9f3c-4d48-b8c2-07d7173f8a32",
						"Title":   "Todoのテストを行う4",
						"Content": "Todo機能のテストをGo言語で行う4",
						"Status":  "Pending",
					},
					{
						"Id":      "c678ce19-cee1-4dd3-a128-b2312c89f2fa",
						"Title":   "Todoのテストを行う3",
						"Content": "Todo機能のテストをGo言語で行う3",
						"Status":  "Completed",
					},
					{
						"Id":      "d11e3442-64f7-4c5e-8368-8ff1e7ad8437",
						"Title":   "Todoのテストを行う6",
						"Content": "Todo機能のテストをGo言語で行う6",
						"Status":  "Pending",
					},
					{
						"Id":      "fa1bbd50-c4b0-4053-85b2-b1c07c42bb10",
						"Title":   "Todoのテストを行う2",
						"Content": "Todo機能のテストをGo言語で行う2",
						"Status":  "InProgress",
					},
				},
				"totalTasks": "6",
			},
		},
		"正常系_QueryParameter付き": {
			url:          "/v1/tasks/?status=Completed",
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"tasks": []map[string]interface{}{
					{
						"Id":      "c678ce19-cee1-4dd3-a128-b2312c89f2fa",
						"Title":   "Todoのテストを行う3",
						"Content": "Todo機能のテストをGo言語で行う3",
						"Status":  "Completed",
					},
				},
				"totalTasks": "1",
			},
		},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			var responseBody map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
				t.Fatalf("failed to unmarshal response body: %v", err)
			}

			if tasks, ok := responseBody["tasks"].([]interface{}); ok {
				var convertedTasks []map[string]interface{}
				for _, task := range tasks {
					if taskMap, ok := task.(map[string]interface{}); ok {
						convertedTasks = append(convertedTasks, taskMap)
					}
				}
				responseBody["tasks"] = convertedTasks
			}

			if diff := cmp.Diff(tt.expectedBody, responseBody); diff != "" {
				t.Errorf("response body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
