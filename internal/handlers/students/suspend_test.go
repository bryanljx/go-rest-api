package students_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bryanljx/go-rest-api/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSuspendHandler(t *testing.T) {
	tests := []struct {
		name    string
		payload map[string]any
		status  int
	}{
		{
			name: "Suspend students",
			payload: map[string]any{
				"student": "studentmary@gmail.com",
			},
			status: http.StatusNoContent,
		},
		{
			name:    "Empty Payload",
			payload: map[string]any{},
			status:  http.StatusBadRequest,
		},
	}

	sh := mocks.MockStudentHandler()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			jsonPayload, _ := json.Marshal(tt.payload)
			payloadReader := bytes.NewReader(jsonPayload)

			r, err := http.NewRequest(http.MethodPost, "/", payloadReader)
			if err != nil {
				t.Error(err)
			}

			sh.Suspend(rr, r)

			rs := rr.Result()

			assert.Equal(t, tt.status, rs.StatusCode)
		})
	}
}
