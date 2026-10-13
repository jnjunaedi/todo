package mysql

import (
	"context"
	"database/sql"
	"errors"
	"todo-app/internal/apperror"
	"todo-app/internal/domain"

	"github.com/google/uuid"
)

type TaskRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTaskRepo(ctx context.Context, task domain.TaskDomain) error {
	query := `INSERT INTO tasks (id, title, description) VALUES (?, ?, ?)`

	genUUID, err := uuid.NewV7()
	if err != nil {
		return apperror.New(apperror.ErrInternalServer, "terjadi kesalahan server")
	}

	_, err = r.db.ExecContext(ctx, query, genUUID.String(), task.Title, task.Description)

	if err != nil {
		return apperror.New(apperror.ErrInternalServer, "terjadi kesalahan server")
	}

	return nil
}

func (r *TaskRepository) FindTaskByIDRepo(ctx context.Context, id string) (domain.TaskDomain, error) {
	query := `SELECT (id, title, description, created_at, updated_at) FROM tasks WHERE id = ?`

	var task domain.TaskDomain

	err := r.db.QueryRowContext(ctx, query, id).Scan()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, apperror.New(apperror.ErrBadRequest, "data tidak ditemukan dari ID "+id)
		}
		return task, apperror.New(apperror.ErrInternalServer, "terjadi kesalahan server")
	}

	return task, nil

}
