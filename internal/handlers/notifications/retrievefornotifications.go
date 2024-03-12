package notifications

import (
	"net/http"
	"strings"

	"github.com/bryanljx/go-rest-api/internal/errorresponse"
	"github.com/bryanljx/go-rest-api/internal/lib/httputils"
	"github.com/bryanljx/go-rest-api/internal/lib/utils"
	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/bryanljx/go-rest-api/internal/models"
)

type RetrieveForNotificationRequest struct {
	Teacher      string `json:"teacher" validate:"required,email"`
	Notification string `json:"notification" validate:"required"`
}

type RetrieveForNotificationResponse struct {
	Recipients []string `json:"recipients"`
}

const (
	ErrValidatingParams = "invalid json body"
)

// @Summary Retrieve recipients (students) for notifications
// @Tags Notification
// @Accept */*
// @Produce json
// @Param RetrieveForNotification body RetrieveForNotificationRequest true
// @Success 200 {object} RetrieveForNotificationResponse
// @Router /retrievefornotification [post]
func (nh *NotificationHandler) RetrieveForNotification(w http.ResponseWriter, r *http.Request) {
	var req RetrieveForNotificationRequest
	ctx := r.Context()

	err := httputils.DecodeJson(r.Body, &req)
	if err != nil {
		httputils.EncodeError(w, nh.logger, errorresponse.BadRequest(err))
		return
	}

	err = nh.Validator.ValidateStruct(req)
	if err != nil {
		httputils.EncodeError(w, nh.logger, errorresponse.ValidationFailed)
		return
	}

	registeredStudentEmails, err := nh.teacherRepo.RetrieveRegisteredStudents(ctx, req.Teacher)
	if err != nil {
		httputils.EncodeError(w, nh.logger, errorresponse.WrapError(err))
		return
	}

	suspendedStudents, err := nh.studentRepo.ListSuspendedStudents(ctx)

	mentionedStudents := extractMentionedStudents(req, nh.Validator, suspendedStudents)

	recipients := utils.MergeSlice[string](mentionedStudents, registeredStudentEmails)

	payload := RetrieveForNotificationResponse{
		Recipients: recipients,
	}

	err = httputils.EncodeJson(w, http.StatusOK, payload, nil)
	if err != nil {
		httputils.EncodeError(w, nh.logger, errorresponse.WrapError(err))
		return
	}
}

func extractMentionedStudents(req RetrieveForNotificationRequest, validator validator.Validator, suspendedStudents []models.Student) []string {
	notif := req.Notification
	tokens := strings.Split(notif, " ")
	mentioned := make([]string, 0, len(tokens))

	for _, t := range tokens {
		if strings.HasPrefix(t, "@") {
			if err := validator.ValidateField(t[1:], "email"); err == nil {
				mentioned = append(mentioned, t[1:])
			}
		}
	}

	res := make([]string, 0, len(mentioned))

	for _, m := range mentioned {
		is_suspended := false
		for _, s := range suspendedStudents {
			if m == s.Email {
				is_suspended = true
				break
			}
		}
		if !is_suspended {
			res = append(res, m)
		}
	}
	return res
}
