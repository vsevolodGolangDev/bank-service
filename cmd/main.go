package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vsevolodGolangDev/bank-service/internal/handler"
	"github.com/vsevolodGolangDev/bank-service/internal/repo"
	"github.com/vsevolodGolangDev/bank-service/internal/service"
	"github.com/vsevolodGolangDev/bank-service/pkg/logging"
)

// @title Balance Service
// @version 1.0
// @description Service for working with user`s balance

// @host localhost:8080
// @BasePath /

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	host := viper.GetString("pg.host")
	if os.Getenv("COMPOSE") == "true" {
		log.Println("Running in compose, using pg.compose_host")
		host = viper.GetString("pg.compose_host")
	} else {
		log.Println("Running in dev, using pg.host")
	}

	// Loading environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to init .env")
	}

	logger, err := logging.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	cfg := repo.Config{
		Port:     viper.GetString("pg.port"),
		Password: os.Getenv("PG_PASSWORD"),
		Username: viper.GetString("pg.username"),
		Host:     host,
		Name:     viper.GetString("pg.name"),
		SSL:      viper.GetString("pg.sslmode"),
	}

	pq, err := repo.InitDB(cfg)
	if err != nil {
		logger.Fatal(err.Error())
	}

	repo := repo.NewRepo(pq, logger)
	service := service.NewService(repo, logger)
	handler := handler.NewHandler(service, logger)

	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), handler.InitRoutes()))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
