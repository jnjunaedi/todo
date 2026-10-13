package domain

import (
	"context"
	"time"
)

type TaskDomain struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskRepository interface {
	CreateTaskRepo(ctx context.Context, task TaskDomain) error
	FindTaskByIDRepo(ctx context.Context, id string) (TaskDomain, error)
}
type TaskUsecase interface {
	CreateTaskUc(ctx context.Context, task TaskDomain) error
	FindTaskByIDUc(ctx context.Context, id string) (TaskDomain, error)
}
