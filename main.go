package main

import (
	"CRUD-REDIS/config"
	"CRUD-REDIS/controller"
	"CRUD-REDIS/database"
	"CRUD-REDIS/repo"
	router2 "CRUD-REDIS/router"
	"CRUD-REDIS/usecase"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	db := database.PostgresConnect(&config)
	rdb := database.RedisConnect(&config)

	startServer(db, rdb)
}

func startServer(db *sql.DB, rdb *redis.Client) {
	app := fiber.New()
	novelRepo := repo.NewNovelRepo(db, rdb)
	noveluseCase := usecase.NewNovelUsecase(novelRepo)
	novelController := controller.NewNovelController(noveluseCase)
	router := router2.NewRouter(app, novelController)
	err := router.Listen(":8085")
	if err != nil {
		log.Fatal(err)
	}
}
