package handler

import (
	"context"
	"net/http"
	"time"
	"todo-app/internal/delivery/http/errhandler"
	"todo-app/internal/delivery/http/request"
	"todo-app/internal/domain"
	"todo-app/internal/pkg/utils"
	"todo-app/internal/pkg/validator"
)

type TaskHandler struct {
	uc domain.TaskUsecase
}

func NewHandler(uc domain.TaskUsecase) *TaskHandler {
	return &TaskHandler{uc: uc}
}

func (h *TaskHandler) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var req request.TaskRequest

	if err := utils.ReadJSON(w, r, &req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Format JSON tidak valid"})
		return
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"status":  http.StatusBadRequest,
			"message": "validasi payload gagal",
			"errors":  errs,
		})
		return
	}

	taskDomain := req.ToDomain()

	if err := h.uc.CreateTaskUc(ctx, taskDomain); err != nil {
		errhandler.HandleInternalError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Catatan berhasil disimpan"})
}

func (h *TaskHandler) HandlerFindTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	path := r.PathValue("id")

	if path == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "id tidak boleh kosong")
		return
	}

	row, err := h.uc.FindTaskByIDUc(ctx, path)

	if err != nil {
		errhandler.HandleInternalError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, row)
}
