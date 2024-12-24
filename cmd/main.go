package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"url-shortener-api/internal/cache"
	"url-shortener-api/internal/controller"
	"url-shortener-api/internal/repository"
	"url-shortener-api/internal/service"
	"url-shortener-api/internal/storage"
)

func main() {
	ctx := context.Background()
	r := gin.Default()

	godotenv.Load()

	redisClient := storage.NewRedisClient(ctx)
	redisCache := cache.NewCache(redisClient)

	databaseConnection, err := storage.NewPgxConnection(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}

	urlRepository := repository.NewURLRepository(databaseConnection)
	urlService := service.NewURLService(urlRepository, redisCache)

	urlController := controller.NewUrlController(urlService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.POST("/urls", urlController.Insert)
	r.GET("/:hash", urlController.Get)

	err = r.Run()
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
