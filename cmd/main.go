package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database"
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database/postgres"
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/handler"
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/service"
	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresConnection(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := database.NewDatabase(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(structure.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("UserSegmentationApp STARTED")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("UserSegmentationApp is SHUTTING DOWN")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured while stopping http server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
