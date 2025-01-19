package main

import (
	"os"
	"praktik-todo/config"
	"praktik-todo/internal/entity"
	"praktik-todo/internal/handler"
	"praktik-todo/internal/repository"
	"praktik-todo/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	/*
		- Log as Json
		- Output to stdout instead of stderr
		- Only log the info or above
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&entity.Task{})
	config.DB.AutoMigrate(&entity.User{})

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Repositories
	taskRepo := repository.NewTaskRepository(config.DB)
	userRepo := repository.NewUserRepository(config.DB)

	// Initialize Usecase
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize Handler
	taskHandler := handler.NewTaskHandler(taskUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	// routes
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/tasks", taskHandler.GetAllTasks)
		api.GET("/tasks/:id", taskHandler.GetTaskByID)
		api.POST("/tasks", taskHandler.CreateTask)
		api.PUT("/tasks/:id", taskHandler.UpdateTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)

		api.GET("/users", userHandler.GetAllUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
		api.POST("/users", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// run
	port := os.Getenv("PORT")
	r.Run(":" + port)
	log.Infof("Running on port %v", port)
}
