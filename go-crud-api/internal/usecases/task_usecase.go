package usecases

import (
	"context"
	"go-crud-api/internal/entities"
	"go-crud-api/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error)
	GetTasks(ctx context.Context) ([]entities.Task, error)
	UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error
	DeleteTask(ctx context.Context, id primitive.ObjectID) error
}

type taskUseCase struct {
	repo repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) TaskUseCase {
	return &taskUseCase{repo: repo}
}

func (u *taskUseCase) CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error) {
	return u.repo.Create(ctx, task)
}

func (u *taskUseCase) GetTasks(ctx context.Context) ([]entities.Task, error) {
	return u.repo.GetAll(ctx)
}

func (u *taskUseCase) UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error {
	return u.repo.Update(ctx, id, task)
}

func (u *taskUseCase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	return u.repo.Delete(ctx, id)
}