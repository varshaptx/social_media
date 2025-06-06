package main

import (
	"fmt"
	"log"
	"social_media/domain/service"
	"social_media/domain/usecase"
	"social_media/handlers"
	"social_media/infrastructure/repository"
	"social_media/router"
)

func main() {
	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	userUseCase := usecase.NewUserUseCase(userService)

	userHandler := handlers.NewUserHandler(userUseCase)

	r := router.SetupRouter(userHandler)

	fmt.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
