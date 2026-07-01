package main

import (
	"backend-smartcost/src/config"
	"backend-smartcost/src/helper"
	"backend-smartcost/src/middleware"
	authController "backend-smartcost/src/modules/auth/controller"
	authHandler "backend-smartcost/src/modules/auth/handler"
	authRouter "backend-smartcost/src/modules/auth/router"
	categoriesController "backend-smartcost/src/modules/categories/controller"
	categoriesHandler "backend-smartcost/src/modules/categories/handler"
	categoriesRouter "backend-smartcost/src/modules/categories/router"
	productsController "backend-smartcost/src/modules/products/controller"
	productsHandler "backend-smartcost/src/modules/products/handler"
	productsRouter "backend-smartcost/src/modules/products/router"
	reportsController "backend-smartcost/src/modules/reports/controller"
	reportsHandler "backend-smartcost/src/modules/reports/handler"
	reportsRouter "backend-smartcost/src/modules/reports/router"
	transactionsController "backend-smartcost/src/modules/transactions/controller"
	transactionsHandler "backend-smartcost/src/modules/transactions/handler"
	transactionsRouter "backend-smartcost/src/modules/transactions/router"
	usersController "backend-smartcost/src/modules/users/controller"
	usersHandler "backend-smartcost/src/modules/users/handler"
	usersRouter "backend-smartcost/src/modules/users/router"
	voidLogsController "backend-smartcost/src/modules/void-logs/controller"
	voidLogsHandler "backend-smartcost/src/modules/void-logs/handler"
	voidLogsRouter "backend-smartcost/src/modules/void-logs/router"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	storage_go "github.com/supabase-community/storage-go"
)

func main() {
	ctx := context.Background()

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "development"
	}

	var envFile string
	if appEnv == "production" {
		envFile = ".env.prod"
	} else {
		envFile = ".env.dev"
	}

	if err := godotenv.Overload(envFile); err != nil {
		log.Printf("Info: File %s not found; using the OS environment variable system\n", envFile)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed load configuration: %v", err)
	}

	// open connection database
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open(cfg.Database.Drive, dsn)
	if err != nil {
		log.Fatalf("Failed to open a database connection: %v", err)
	}

	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.Database.ConnMaxIdleTime) * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("Database not responding: %v", err)
	}

	defer db.Close()

	log.Printf("Successfully connected to the database using the driver: %s", cfg.Database.Drive)

	// open connection redis
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB: cfg.Redis.DB,
		PoolSize: cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
		DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed connected a redis: %v", err)
	}

	defer rdb.Close()

	log.Println("Successfully connected to the redis")

	// storage client
	storageClient := storage_go.NewClient(cfg.Supabase.URL, cfg.Supabase.SecretAccessKey, nil)

	// init helper
	helper := helper.NewHelper(cfg)

	// init middleware
	middleware := middleware.NewMiddleware(cfg, helper, rdb)

	// setup gin
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.SecureHeaders()) // Proteksi Security Header
	r.Use(middleware.Timeout(time.Duration(cfg.App.TimeoutSeconds) * time.Second)) // Request Timeout

	api := r.Group("/api/v1")

	// health check endpoint
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      true,
			"message":     "Server is healthy",
			"environment": cfg.App.Env,
			"timestamp":   time.Now().Format(time.RFC3339),
		})
	})

	// Handle 404 Not Found standar API profesional
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Endpoint not found"})
	})

	// init modules

	// auth
	authController := authController.NewController(db, rdb, storageClient, helper)
	authHandler := authHandler.NewHandler(authController, helper)
	authRouter.New(api, middleware, authHandler)

	// users
	usersController := usersController.NewController(db, rdb, storageClient, helper)
	usersHandler := usersHandler.NewHandler(usersController, helper)
	usersRouter.New(api, middleware, usersHandler)

	// products
	productsController := productsController.NewController(db, rdb, storageClient, helper)
	productsHandler := productsHandler.NewHandler(productsController, helper)
	productsRouter.New(api, middleware, productsHandler)

	// categories
	categoriesController := categoriesController.NewController(db, rdb, storageClient, helper)
	categoriesHandler := categoriesHandler.NewHandler(categoriesController, helper)
	categoriesRouter.New(api, middleware, categoriesHandler)

	// reports
	reportsController := reportsController.NewController(db, rdb, storageClient, helper)
	reportsHandler := reportsHandler.NewHandler(reportsController, helper)
	reportsRouter.New(api, middleware, reportsHandler)

	// transactions
	transactionsController := transactionsController.NewController(db, rdb, storageClient, helper)
	transactionsHandler := transactionsHandler.NewHandler(transactionsController, helper)
	transactionsRouter.New(api, middleware, transactionsHandler)

	// void-logs
	voidLogsController := voidLogsController.NewController(db, rdb, storageClient, helper)
	voidLogsHandler := voidLogsHandler.NewHandler(voidLogsController, helper)
	voidLogsRouter.New(api, middleware, voidLogsHandler)

	// running server
	log.Printf("server running on port %s", cfg.App.Port)
	if err := r.Run(fmt.Sprintf("%s%s", ":", cfg.App.Port)); err != nil {
		log.Println("failed to running server")
	}
}