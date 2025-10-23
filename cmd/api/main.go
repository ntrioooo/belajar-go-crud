package main

import (
	"log"
	"time"

	"belajar-go/internal/config"
	"belajar-go/internal/core/services"
	"belajar-go/internal/http/handlers"
	"belajar-go/internal/http/middleware"
	"belajar-go/internal/repository"
	gormrepo "belajar-go/internal/repository/gorm"
	"belajar-go/pkg/jwtutil"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := repository.Open(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate (opsional, bisa via migrations tool)
	if err := db.AutoMigrate(
		&gormrepo.User{},
		&gormrepo.Post{},
		&gormrepo.Like{},
		&gormrepo.Category{},
		&gormrepo.Comment{},
		&gormrepo.CommentLike{},
		&gormrepo.Retweet{}); err != nil {
		log.Fatal(err)
	}

	// Wiring DI
	userRepo := gormrepo.NewUserRepository(db)
	postRepo := gormrepo.NewPostRepository(db)
	categoryRepo := gormrepo.NewCategoryRepository(db)
	commentRepo := gormrepo.NewCommentRepository(db)
	retweetRepo := gormrepo.NewRetweetRepository(db)

	jwtm := jwtutil.New(cfg.Secret)

	authSvc := services.NewAuthService(userRepo, jwtm)
	postSvc := services.NewPostService(postRepo)
	categorySvc := services.NewCategoryService(categoryRepo)
	userSvc := services.NewUserService(userRepo)
	commentSvc := services.NewCommentService(commentRepo, postRepo)
	retweetSvc := services.NewRetweetService(retweetRepo, postRepo)
	profileSvc := services.NewProfileService(userRepo, postRepo, retweetRepo, commentRepo)

	userH := handlers.NewUserHandler(userSvc)
	authH := handlers.NewAuthHandler(authSvc)
	postH := handlers.NewPostHandler(postSvc)
	categoryH := handlers.NewCategoryHandler(categorySvc)
	commentH := handlers.NewCommentHandler(commentSvc)
	retweetH := handlers.NewRetweetHandler(retweetSvc)
	profileH := handlers.NewProfileHandler(profileSvc, userRepo)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"}, // JANGAN pakai "*"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
		v1.POST("/categories", middleware.RequireAuth(jwtm), middleware.RequireAdmin(userRepo), categoryH.Create)
		v1.PUT("/categories/:id", middleware.RequireAuth(jwtm), middleware.RequireAdmin(userRepo), categoryH.Update)
		v1.DELETE("/categories/:id", middleware.RequireAuth(jwtm), middleware.RequireAdmin(userRepo), categoryH.Delete)

		// comments
		v1.POST("/posts/:id/comments", middleware.RequireAuth(jwtm), commentH.Create)
		v1.GET("/posts/:id/comments", middleware.OptionalAuth(jwtm), commentH.ListByPost)
		v1.GET("/comments/:id/replies", middleware.OptionalAuth(jwtm), commentH.ListReplies)
		v1.POST("/comments/:id/like", middleware.RequireAuth(jwtm), commentH.ToggleLike)
		v1.DELETE("/comments/:id", middleware.RequireAuth(jwtm), commentH.Delete)

		// retweet / quote
		v1.POST("/posts/:id/retweet", middleware.RequireAuth(jwtm), retweetH.ToggleRetweet)
		v1.POST("/posts/:id/quote", middleware.RequireAuth(jwtm), retweetH.ToggleQuote)

		// profile
		v1.GET("/profiles/:username", middleware.OptionalAuth(jwtm), profileH.GetByUsername)
	}

	log.Println("listening on :" + cfg.Port)
	r.Run(":" + cfg.Port)
}
