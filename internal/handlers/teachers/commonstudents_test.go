package teachers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bryanljx/go-rest-api/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCommonStudents(t *testing.T) {
	tests := []struct {
		name         string
		urlParam     string
		expectedResp map[string]any
		status       int
	}{
		{
			name:     "Get common students between teachers",
			urlParam: "?teacher=random@gmail.com",
			expectedResp: map[string]any{
				"students": []string{
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
			name:     "Missing param",
			urlParam: "",
			expectedResp: map[string]any{
				"message": "missing url param - teacher",
			},
			status: http.StatusBadRequest,
		},
	}

	th := mocks.MockTeacherHandler()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			r, err := http.NewRequest(http.MethodPost, "/api/commonstudents"+tt.urlParam, nil)
			if err != nil {
				t.Error(err)
			}

			th.CommonStudents(rr, r)

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
