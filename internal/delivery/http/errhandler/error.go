package errhandler

import (
	"errors"
	"net/http"
	"todo-app/internal/apperror"
	"todo-app/internal/pkg/utils"
)

func HandleInternalError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	message := "Terjadi kesalahan pada server internal"

	switch {
	case errors.Is(err, apperror.ErrBadRequest):
		statusCode = http.StatusBadRequest
		message = err.Error()
	case errors.Is(err, apperror.ErrNotFound):
		statusCode = http.StatusNotFound
		message = err.Error()
	case errors.Is(err, apperror.ErrInternalServer):
		statusCode = http.StatusInternalServerError
		message = err.Error()
	}

	utils.WriteJSON(w, statusCode, map[string]any{
		"status": statusCode,
		"error":  message,
	})
}
