package main

import (
	"context"
	"dev11"
	"dev11/pkg/handler"
	"dev11/pkg/handler/middleware"
	"dev11/pkg/repository"
	"dev11/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}

	repos := repository.NewStorage()
	services := service.NewCalendarService(repos)
	loggers := middleware.NewLogger("log.txt", "info")
	handlers := handler.NewHandler(services, repos, loggers)
	router := handlers.InitRoutes()

	router.Use(middleware.LoggerMiddleware(loggers))

	srv := new(dev11.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), router); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("CalendarApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("CalendarApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
