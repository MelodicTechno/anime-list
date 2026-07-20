package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/MelodicTechno/anime-list/internal/config"
	"github.com/MelodicTechno/anime-list/internal/database"
	"github.com/MelodicTechno/anime-list/internal/handler"
	"github.com/MelodicTechno/anime-list/internal/repository"
	"github.com/MelodicTechno/anime-list/internal/service"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.NewPostgres(&cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	log.Println("connected to postgres")

	rdb, err := database.NewRedis(&cfg.Redis)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("connected to redis")

	_ = db

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, rdb, cfg.JWT.Secret, cfg.JWT.AccessExpireHours, cfg.JWT.RefreshExpireHours)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.POST("/refresh", userHandler.Refresh)

		auth := api.Group("")
		auth.Use(handler.AuthMiddleware(cfg.JWT.Secret))
		{
			auth.GET("/me", userHandler.Me)
		}
	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
