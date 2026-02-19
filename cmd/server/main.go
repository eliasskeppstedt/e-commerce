package main

import (
	"database/sql"
	"ecommerce/duckyarmy/api"
	"ecommerce/duckyarmy/internal/customer"
	"ecommerce/duckyarmy/internal/product"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// Create Gin router
	engine := gin.Default()

	// Initialize database
	db := tmpDbConfig()

	// USER SETUP
	userRepo := customer.NewMysqlUserRepository(db)
	userService := customer.NewUserService1(userRepo)
	userHandler := customer.NewUserHandler(userService)

	// PRODUCT SETUP
	productRepo := product.NewMysqlProductRepository(db)
	productService := product.NewProductServiceImp(productRepo)
	productHandler := product.NewProductHandler(productService)

	// Load HTML and static files
	engine.LoadHTMLGlob("web/html/*")
	engine.Static("/styles", "./web/styles")
	engine.Static("/icons", "./web/icons")

	// Register routes
	api.RegisterWebRouts(engine)
	api.RegisterApiRouts(engine, userHandler, productHandler)

	// Start server
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func tmpDbConfig() *sql.DB {

	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// MySQL config
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBURL")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.ParseTime = true

	// Open connection
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("\n-- Connected to database successfully --")

	return db
}
