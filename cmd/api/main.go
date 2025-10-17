package main

import (
	"log"

	"belajar-go/internal/config"
	"belajar-go/internal/core/services"
	"belajar-go/internal/http/handlers"
	"belajar-go/internal/http/middleware"
	"belajar-go/internal/repository"
	gormrepo "belajar-go/internal/repository/gorm"
	"belajar-go/pkg/jwtutil"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := repository.Open(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate (opsional, bisa via migrations tool)
	if err := db.AutoMigrate(&gormrepo.User{}, &gormrepo.Post{}); err != nil {
		log.Fatal(err)
	}

	// Wiring DI
	userRepo := gormrepo.NewUserRepository(db)
	postRepo := gormrepo.NewPostRepository(db)

	jwtm := jwtutil.New(cfg.Secret)

	authSvc := services.NewAuthService(userRepo, jwtm)
	postSvc := services.NewPostService(postRepo)

	authH := handlers.NewAuthHandler(authSvc)
	postH := handlers.NewPostHandler(postSvc)

	r := gin.Default()
	r.Use(middleware.RequestID())

	v1 := r.Group("/v1")
	{
		v1.POST("/auth/signup", authH.Signup)
		v1.POST("/auth/login", authH.Login)

		v1.GET("/posts", postH.List)
		v1.GET("/posts/:id", postH.Show)
		v1.POST("/posts", middleware.RequireAuth(jwtm), postH.Create) // bisa tambahkan RequireAuth di sini
		v1.PUT("/posts/:id", middleware.RequireAuth(jwtm), postH.Update)
		v1.DELETE("/posts/:id", middleware.RequireAuth(jwtm), postH.Delete)
	}

	log.Println("listening on :" + cfg.Port)
	r.Run(":" + cfg.Port)
}
