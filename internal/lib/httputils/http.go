package httputils

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/bryanljx/go-rest-api/internal/errorresponse"
)

func EncodeJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}

func EncodeError(w http.ResponseWriter, logger *slog.Logger, err errorresponse.ErrorResponse) {
	encodeErr := EncodeJson(w, err.HttpStatusCode, err, nil)
	if encodeErr != nil {
		logger.Error("failed to write error message")
	}
}

func DecodeJson(r io.Reader, view any) error {
	decoder := json.NewDecoder(r)
	// TODO: Consider enabling if frontend is fine with this
	// decoder.DisallowUnknownFields()
	return decoder.Decode(view)
}

func WriteRespNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
