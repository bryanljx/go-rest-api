package teachers

import (
	"net/http"

	"github.com/bryanljx/go-rest-api/internal/errorresponse"
	"github.com/bryanljx/go-rest-api/internal/lib/httputils"
)

const (
	teacherParam = "teacher"

	MissingTeacherParameterErr = "missing url param - teacher"
)

type CommonStudentsResponse struct {
	Students []string `json:"students"`
}

// @Summary Retrieve common students
// @Tags Teacher
// @Accept */*
// @Produce json
// @Param teachers url string true "List of teachers' emails"
// @Success 200 {object} CommonStudentsResponse
// @Router /commonstudents [post]
func (th *TeacherHandler) CommonStudents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := r.URL.Query()
	if !params.Has(teacherParam) {
		httputils.EncodeError(w, th.logger, errorresponse.BadRequest(
			errorresponse.NewError(MissingTeacherParameterErr, http.StatusBadRequest),
		))
		return
	}

	teachers := params[teacherParam]

	err := th.Validator.ValidateField(teachers, "dive,email")
	if err != nil {
		httputils.EncodeError(w, th.logger, errorresponse.ValidationFailed)
		return
	}

	studentEmails, err := th.repo.CommonStudents(ctx, teachers)
	if err != nil {
		httputils.EncodeError(w, th.logger, errorresponse.WrapError(err))
		return
	}

	payload := CommonStudentsResponse{
		Students: studentEmails,
	}

	err = httputils.EncodeJson(w, http.StatusOK, payload, nil)
	if err != nil {
		httputils.EncodeError(w, th.logger, errorresponse.WrapError(err))
		return
	}
}
