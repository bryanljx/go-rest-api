package teachers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bryanljx/go-rest-api/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveForNotificationHandler(t *testing.T) {
	tests := []struct {
		name         string
		payload      map[string]any
		expectedResp map[string]any
		status       int
	}{
		{
			name: "Register students",
			payload: map[string]any{
				"teacher": "teacherken@gmail.com",
				"students": []string{
					"studentjon@gmail.com",
					"studenthon@gmail.com",
				},
			},
			status: http.StatusNoContent,
		},
		{
			name:    "Empty Payload",
			payload: map[string]any{},
			status:  http.StatusBadRequest,
		},
	}

	th := mocks.MockTeacherHandler()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			jsonPayload, _ := json.Marshal(tt.payload)
			payloadReader := bytes.NewReader(jsonPayload)

			r, err := http.NewRequest(http.MethodPost, "/", payloadReader)
			if err != nil {
				t.Error(err)
			}

			th.RegisterStudents(rr, r)

			rs := rr.Result()

			assert.Equal(t, tt.status, rs.StatusCode)
		})
	}
}
