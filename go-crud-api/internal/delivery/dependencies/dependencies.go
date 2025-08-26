package dependencies

import (
	"go-crud-api/internal/infra"
	"go-crud-api/internal/interfaces/handlers"
	"go-crud-api/internal/repositories"
	"go-crud-api/internal/usecases"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func Setup() *dig.Container {
	container := dig.New()

	if err := container.Provide(infra.NewMongoDatabase); err != nil {
		log.Fatalf("Failed to provide MongoDatabase: %v", err)
	}

	if err := container.Provide(func(db *mongo.Database) repositories.TaskRepository {
		return repositories.NewTaskRepository(db)
	}); err != nil {
		log.Fatalf("Failed to provide TaskRepository: %v", err)
	}

	if err := container.Provide(usecases.NewTaskUseCase); err != nil {
		log.Fatalf("Failed to provide TaskUseCase: %v", err)
	}

	if err := container.Provide(handlers.NewTaskHandler); err != nil {
		log.Fatalf("Failed to provide TaskHandler: %v", err)
	}

	return container
}
