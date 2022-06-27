package main

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/config"
	"github.com/Thunderbirrd/ShortLinksService/internal/handler"
	"github.com/Thunderbirrd/ShortLinksService/internal/repository"
	"github.com/Thunderbirrd/ShortLinksService/internal/repository/postgres"
	"github.com/Thunderbirrd/ShortLinksService/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg, err := config.New()
	if err != nil {
		logrus.Errorf("Error initializing jwtConfig: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(*cfg)
	if err != nil {
		logrus.Fatalf("Failed to initialed db: %s", err.Error())
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var repos *repository.Repository
	if cfg.Mode == "postgres" {
		repos = repository.NewRepositoryPostgres(db)
	} else if cfg.Mode == "internal" {
		repos = repository.NewRepositoryInternal(make(map[string]string))
	} else {
		panic("Wrong Mode value")
	}

	services := service.NewService(repos)
	handlers := handler.NewHandler(services, cfg.HttpPort)

	e.POST("/create", handlers.CreateShortUrl)
	e.GET("/:shortUrl", handlers.RedirectToLongUrl)
	e.GET("/url/:shortUrl", handlers.GetLongUrl)

	e.Logger.Fatal(e.Start(cfg.HttpPort))
}
