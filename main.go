package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kishima/golang-realworld-example-app/config"
	"github.com/k0kishima/golang-realworld-example-app/db"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/handlers"
	"github.com/k0kishima/golang-realworld-example-app/middlewares"
)

func main() {
	config.LoadEnv()

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	dataSourceName := db.GetDataSourceName()
	client, err := ent.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	api := r.Group("/api")
	{
		api.POST("/users", handlers.RegisterUser(client))
		api.POST("/users/login", handlers.Login(client))
		api.GET("/profiles/:username", handlers.GetProfile(client))
		api.GET("/articles", handlers.ListArticles(client))
		api.GET("/articles/:slug", handlers.GetArticle(client))
		api.GET("/articles/:slug/comments", handlers.GetComments(client))
		api.GET("/tags", handlers.GetTags(client))

		api.Use(middlewares.AuthMiddleware(client))
		{
			api.GET("/user", handlers.GetCurrentUser(client))
			api.PUT("/user", handlers.UpdateUser(client))
			api.POST("/profiles/:username/follow", handlers.FollowUser(client))
			api.DELETE("/profiles/:username/follow", handlers.UnfollowUser(client))
			api.POST("/articles", handlers.CreateArticle(client))
			api.PUT("/articles/:slug", handlers.UpdateArticle(client))
			api.DELETE("/articles/:slug", handlers.DeleteArticle(client))
			api.POST("/articles/:slug/comments", handlers.PostComment(client))
			api.DELETE("/articles/:slug/comments/:id", handlers.DeleteComment(client))
			api.POST("/articles/:slug/favorite", handlers.FavoriteArticle(client))
			api.DELETE("/articles/:slug/favorite", handlers.UnfavoriteArticle(client))
			api.GET("/articles/feed", handlers.GetFeed(client))
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server: ", err)
	}
}
