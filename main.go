package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/challenge_prueba_biblioteca/src/infrastructure/mongo"
	"github.com/challenge_prueba_biblioteca/src/infrastructure/web"
	handler "github.com/challenge_prueba_biblioteca/src/infrastructure/web/hanlder"
	"github.com/challenge_prueba_biblioteca/src/interface/repository"
	"github.com/challenge_prueba_biblioteca/src/shared"
	"github.com/challenge_prueba_biblioteca/src/usecase"
)

var logger = log.WithFields(log.Fields{
	"layer": shared.MainLayer,
})

func init() {
	logger.Info("Staring app...")
	godotenv.Load(".env")
}

func main() {
	// app and db instance
	app := web.ServerInstance()
	mongoInstance := mongo.MongoInstance()

	// HEALTH SERVICES
	handler.NewHealthHandler(app, usecase.NewHealthUseCase())

	//Book SERVICES
	bookRepository := repository.NewBookRepository(mongoInstance)
	bookUseCase := usecase.NewBookUseCaseUseCase(bookRepository)
	handler.NewBookHandler(app, bookUseCase)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}
