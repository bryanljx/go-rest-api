package teachers

import (
	"net/http"

	"github.com/bryanljx/go-rest-api/internal/errorresponse"
	"github.com/bryanljx/go-rest-api/internal/lib/httputils"
)

type RegisterStudentsRequest struct {
	Teacher  string   `json:"teacher" validate:"required,email"`
	Students []string `json:"students" validate:"dive,required,email"`
}

// @Summary Register students
// @Tags Teacher
// @Accept */*
// @Produce json
// @Param RegisterStudentsRequest body RegisterStudentsRequest true
// @Success 204
// @Router /register [post]
func (th *TeacherHandler) RegisterStudents(w http.ResponseWriter, r *http.Request) {
	var req RegisterStudentsRequest
	ctx := r.Context()

	err := httputils.DecodeJson(r.Body, &req)
	if err != nil {
		httputils.EncodeError(w, th.logger, errorresponse.BadRequest(err))
		return
	}

	err = th.Validator.ValidateStruct(req)
	if err != nil {
		httputils.EncodeError(w, th.logger, errorresponse.ValidationFailed)
		return
	}

	err = th.repo.RegisterStudents(ctx, req.Teacher, req.Students)
	if err != nil {
		httputils.EncodeError(w, th.logger, errorresponse.WrapError(err))
		return
	}

	httputils.WriteRespNoContent(w)
}
