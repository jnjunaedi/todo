package usecase

import (
	"context"
	"todo-app/internal/domain"
)

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewUsecase(repo domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (u *TaskUsecase) CreateTaskUc(ctx context.Context, task domain.TaskDomain) error {
	return u.repo.CreateTaskRepo(ctx, task)
}

func (u *TaskUsecase) FindTaskByIDUc(ctx context.Context, id string) (domain.TaskDomain, error) {
	return u.repo.FindTaskByIDRepo(ctx, id)
}
