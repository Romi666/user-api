package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-api/config"
	userHandler "user-api/user/handler"
	userRepo "user-api/user/repo"
	userUsecase "user-api/user/usecase"

	"github.com/joho/godotenv"
)

func main() {
	// load env variable from .env.example file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env.example file")
	}

	db := config.DBInit()
	defer db.Close()

	router := gin.Default()

	userRepo := userRepo.CreateUserRepo(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)
	userHandler.CreatePersonHandler(router, userUsecase)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
