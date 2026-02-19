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

	var db *sql.DB // where to store, in app config, any of its child structs or own struct???

	// Create a Gin router with default middleware (logger and recovery)
	engine := gin.Default()

	tmpDbConfig(db)
	fmt.Println("engine")
	userRepo := customer.NewMysqlUserRepository(db)
	fmt.Println("check repo")
	userService := customer.NewUserService1(userRepo)
	fmt.Println("check service")
	userHandler := customer.NewUserHandler(userService)
	fmt.Println("check handler")

	productRepo := product.NewMysqlProductRepository(db)
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductHandler(productService)

	//Load HTML files and css
	engine.LoadHTMLGlob("web/html/*")
	engine.Static("/styles", "./web/styles")
	engine.Static("/icons", "./web/icons")

	api.RegisterWebRouts(engine)
	api.RegisterApiRouts(engine, userHandler, productHandler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}

func tmpDbConfig(db *sql.DB) {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBURL")
	cfg.DBName = os.Getenv("DBNAME")

	// Get a database handle.

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Printf("\n -- Connected!\n\n")
}
