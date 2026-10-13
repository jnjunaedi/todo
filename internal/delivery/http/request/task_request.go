package request

import (
	"todo-app/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (t *TaskRequest) ToDomain() domain.TaskDomain {
	return domain.TaskDomain{
		Title:       t.Title,
		Description: t.Description,
	}
}
