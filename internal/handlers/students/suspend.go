package students

import (
	"net/http"

	"github.com/bryanljx/go-rest-api/internal/errorresponse"
	"github.com/bryanljx/go-rest-api/internal/lib/httputils"
)

type SuspendRequest struct {
	Student string `json:"student" validate:"required,email"`
}

//     • Endpoint: POST /api/suspend
//     • Headers: Content-Type: application/json
//     • Success response status: HTTP 204
//     • Request body example:
// {
//   "student" : "studentmary@gmail.com"
// }

// @Summary Suspend student
// @Tags Student
// @Accept */*
// @Produce json
// @Param SuspendRequest body SuspendRequest true
// @Success 204
// @Router /suspend [post]
func (sh *StudentHandler) Suspend(w http.ResponseWriter, r *http.Request) {
	var req SuspendRequest
	ctx := r.Context()

	err := httputils.DecodeJson(r.Body, &req)
	if err != nil {
		httputils.EncodeError(w, sh.logger, errorresponse.BadRequest(err))
		return
	}

	err = sh.Validator.ValidateStruct(req)
	if err != nil {
		httputils.EncodeError(w, sh.logger, errorresponse.ValidationFailed)
		return
	}

	err = sh.repo.SuspendStudent(ctx, req.Student)
	if err != nil {
		httputils.EncodeError(w, sh.logger, errorresponse.WrapError(err))
		return
	}

	httputils.WriteRespNoContent(w)
}
