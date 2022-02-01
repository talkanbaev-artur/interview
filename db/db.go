package db

import (
	"fmt"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/talkanbaev-artur/interview/config"
	"github.com/talkanbaev-artur/shutdown"
)

func ConnectRelDatabase(conf config.AppConfig, shutdown *shutdown.Shutdown) (rel.Repository, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	adapter, err := postgres.Open(dsn)
	if err != nil {
		return nil, err
	}
	shutdown.Add(adapter.Close) // add close to the global shutdown
	repo := rel.New(adapter)

	return repo, nil
}
