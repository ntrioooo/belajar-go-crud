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
	if err := db.AutoMigrate(&gormrepo.User{}, &gormrepo.Post{}, &gormrepo.Like{}, &gormrepo.Category{}); err != nil {
		log.Fatal(err)
	}

	// Wiring DI
	userRepo := gormrepo.NewUserRepository(db)
	postRepo := gormrepo.NewPostRepository(db)
	categoryRepo := gormrepo.NewCategoryRepository(db)

	jwtm := jwtutil.New(cfg.Secret)

	authSvc := services.NewAuthService(userRepo, jwtm)
	postSvc := services.NewPostService(postRepo)
	categorySvc := services.NewCategoryService(categoryRepo)
	userSvc := services.NewUserService(userRepo)

	userH := handlers.NewUserHandler(userSvc)
	authH := handlers.NewAuthHandler(authSvc)
	postH := handlers.NewPostHandler(postSvc)
	categoryH := handlers.NewCategoryHandler(categorySvc)

	r := gin.Default()
	r.Use(middleware.RequestID())

	v1 := r.Group("/v1")
	{
		v1.POST("/auth/signup", authH.Signup)
		v1.POST("/auth/login", authH.Login)

		v1.GET("/posts", middleware.OptionalAuth(jwtm), postH.List)
		v1.GET("/posts/:id", middleware.OptionalAuth(jwtm), postH.Show)
		v1.GET("/users/me", middleware.RequireAuth(jwtm), userH.Me)
		v1.PUT("/users/me", middleware.RequireAuth(jwtm), userH.UpdateMe)

		// butuh login untuk create/update/delete
		v1.POST("/posts", middleware.RequireAuth(jwtm), postH.Create)
		v1.PUT("/posts/:id", middleware.RequireAuth(jwtm), postH.Update)
		v1.DELETE("/posts/:id", middleware.RequireAuth(jwtm), postH.Delete)
		v1.POST("/posts/:id/like", middleware.RequireAuth(jwtm), postH.ToggleLike)

		v1.GET("/categories", categoryH.List)
		v1.GET("/categories/:id", categoryH.Show)
		v1.POST("/categories", middleware.RequireAuth(jwtm), categoryH.Create)
		v1.PUT("/categories/:id", middleware.RequireAuth(jwtm), categoryH.Update)
		v1.DELETE("/categories/:id", middleware.RequireAuth(jwtm), categoryH.Delete)
	}

	log.Println("listening on :" + cfg.Port)
	r.Run(":" + cfg.Port)
}
