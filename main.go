package main

import (
	"fmt"
	"os"

	"github.com/challenge_prueba_biblioteca/src/infrastructure/web"
	handler "github.com/challenge_prueba_biblioteca/src/infrastructure/web/hanlder"
	"github.com/challenge_prueba_biblioteca/src/shared"
	"github.com/challenge_prueba_biblioteca/src/usecase"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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

	// HEALTH SERVICES
	handler.NewHealthHandler(app, usecase.NewHealthUseCase())

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}
