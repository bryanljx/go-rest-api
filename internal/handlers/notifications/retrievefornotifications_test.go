package notifications_test

import (
	"bytes"
	"encoding/json"
	"io"
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
			name: "Retrieved registed students",
			payload: map[string]any{
				"teacher":      "teacherken@gmail.com",
				"notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com",
			},
			expectedResp: map[string]any{
				"recipients": []string{
					"studentagnes@gmail.com",
					"studentmiche@gmail.com",
					"commonstudent1@gmail.com",
					"commonstudent2@gmail.com",
					"student_only_under_teacher_ken@gmail.com",
					"studentmary@gmail.com",
					"studentbob@gmail.com",
				},
			},
			status: http.StatusOK,
		},
		{
			name: "Retrieved @mentioned students",
			payload: map[string]any{
				"teacher":      "teacherken@gmail.com",
				"notification": "Hello students! @test@gmail.com",
			},
			expectedResp: map[string]any{
				"recipients": []string{
					"test@gmail.com",
					"commonstudent1@gmail.com",
					"commonstudent2@gmail.com",
					"student_only_under_teacher_ken@gmail.com",
					"studentmary@gmail.com",
					"studentbob@gmail.com",
					"studentagnes@gmail.com",
					"studentmiche@gmail.com",
				},
			},
			status: http.StatusOK,
		},
		{
			name:    "Empty Payload",
			payload: map[string]any{},
			expectedResp: map[string]any{
				"message": "Bad Request - invalid json body",
			},
			status: http.StatusBadRequest,
		},
		{
			name: "SuspendedMentionedStudents",
			payload: map[string]any{
				"teacher":      "teacherken@gmail.com",
				"notification": "Hello students! @JohnDoe@gmail.com @Guy123@gmail.com",
			},
			expectedResp: map[string]any{
				"recipients": []string{
					"commonstudent1@gmail.com",
					"commonstudent2@gmail.com",
					"student_only_under_teacher_ken@gmail.com",
					"studentmary@gmail.com",
					"studentbob@gmail.com",
					"studentagnes@gmail.com",
					"studentmiche@gmail.com",
				},
			},
			status: http.StatusOK,
		},
	}

	nh := mocks.MockNotificationHandler()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			jsonPayload, _ := json.Marshal(tt.payload)
			payloadReader := bytes.NewReader(jsonPayload)

			r, err := http.NewRequest(http.MethodPost, "/", payloadReader)
			if err != nil {
				t.Error(err)
			}

			nh.RetrieveForNotification(rr, r)

			rs := rr.Result()

			assert.Equal(t, tt.status, rs.StatusCode)

			defer rs.Body.Close()
			body, err := io.ReadAll(rs.Body)
			if err != nil {
				t.Error(err)
			}

			expectedBody, err := json.Marshal(tt.expectedResp)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, string(expectedBody), string(body))
		})
	}
}
