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
	tweetRepo := repository.NewTweetRepository(userRepo.GetAllUsers())

	userService := service.NewUserService(userRepo)
	tweetService := service.NewTweetService(tweetRepo, userService)

	userUseCase := usecase.NewUserUseCase(userService, tweetService)

	userHandler := handlers.NewUserHandler(userUseCase)

	r := router.SetupRouter(userHandler)

	fmt.Println("Server starting on :8085")
	if err := r.Run(":8085"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
