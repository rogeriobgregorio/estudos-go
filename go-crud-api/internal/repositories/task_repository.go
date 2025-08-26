package repositories

import (
	"context"
	"go-crud-api/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entities.Task) (primitive.ObjectID, error)
	GetAll(ctx context.Context) ([]entities.Task, error)
	Update(ctx context.Context, id primitive.ObjectID, task *entities.Task) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{
		collection: db.Collection("tasks"),
	}
}

func (r *taskRepository) Create(ctx context.Context, task *entities.Task) (primitive.ObjectID, error) {
	task.CreatedAt = time.Now()
	result, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *taskRepository) GetAll(ctx context.Context) ([]entities.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []entities.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, id primitive.ObjectID, task *entities.Task) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": task})
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}