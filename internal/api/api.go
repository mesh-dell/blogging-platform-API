package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mesh-dell/blogging-platform-API/config"
	"github.com/mesh-dell/blogging-platform-API/internal/blog/handler"
	"github.com/mesh-dell/blogging-platform-API/internal/blog/repository"
	"github.com/mesh-dell/blogging-platform-API/internal/blog/service"
	"github.com/mesh-dell/blogging-platform-API/internal/database"
)

func InitServer(config config.Config) {
	// create new gin app
	router := gin.Default()

	db, err := database.NewMySQLDB(config)
	if err != nil {
		log.Fatalf("init db: %v", err)
	}

	blogRepository := repository.NewBlogRepository(db.Client)
	blogPostService := service.NewBlogPostService(blogRepository)
	blogHandler := handler.NewBlogHandler(*blogPostService)
	// routes
	router.GET("/posts", blogHandler.GetAll)
	router.GET("/posts/:id", blogHandler.GetByID)
	router.POST("/posts", blogHandler.Create)
	router.PUT("/posts/:id", blogHandler.Update)
	router.DELETE("/posts/:id", blogHandler.Delete)
	router.Run(":" + config.Port)
}
