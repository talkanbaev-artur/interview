package main

import (
	"context"
	"log"

	"github.com/talkanbaev-artur/interview/config"
	"github.com/talkanbaev-artur/interview/db"
	"github.com/talkanbaev-artur/interview/user/http"
	"github.com/talkanbaev-artur/interview/user/repo"
	"github.com/talkanbaev-artur/interview/user/service"
	"github.com/talkanbaev-artur/shutdown"
	"go.uber.org/zap"
)

func main() {
	config := config.ReadConfig()
	shutd := shutdown.NewShutdown()
	logger := InitLogger(shutd)
	ctx, cancel := context.WithCancel(context.Background())

	relRepo, err := db.ConnectRelDatabase(config, shutd)
	if err != nil {
		logger.Fatalw("Could not connect to the database", "Error", err)
	}

	//init user
	userRepo := repo.NewRelRepo(relRepo)
	us := service.NewService(userRepo, logger.With("level", "service"))

	go http.NewHTTPServer(cancel, config, us)

	logger.Infow("Initialised authentication server", "ConfigStatus", "OK", "ServiceStatus", "OK")
	shutdown.Wait(ctx, cancel, shutd)
}

func InitLogger(shutdown *shutdown.Shutdown) *zap.SugaredLogger {
	log1, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	logger := log1.Sugar()
	zap.ReplaceGlobals(log1)
	shutdown.Add(logger.Sync)
	return logger
}
