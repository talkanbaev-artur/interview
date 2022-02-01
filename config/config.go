package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
	ServerPort int
}

func ReadConfig() AppConfig {
	//create deafult config
	conf := AppConfig{"localhost", "postgres", "postgres", "interview_test", 5432, 8000}

	if key := os.Getenv("DB_HOST"); key != "" {
		conf.DBHost = key
	}

	if key := os.Getenv("DB_PORT"); key != "" {
		//TODO: handle possible parsing error
		port, _ := strconv.ParseInt(key, 10, 64)
		conf.DBPort = int(port)
	}

	if key := os.Getenv("PORT"); key != "" {
		//TODO: handle possible parsing error
		port, _ := strconv.ParseInt(key, 10, 64)
		conf.ServerPort = int(port)
	}

	if key := os.Getenv("DB_USER"); key != "" {
		conf.DBUser = key
	}

	if key := os.Getenv("DB_PASSWORD"); key != "" {
		conf.DBPassword = key
	}

	if key := os.Getenv("DB_NAME"); key != "" {
		conf.DBName = key
	}

	return conf
}
