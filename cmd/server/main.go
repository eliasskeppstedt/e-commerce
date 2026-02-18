package main

import (
	"database/sql"
	"ecommerce/duckyarmy/api"
	"ecommerce/duckyarmy/internal/customer"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB // where to store, in app config, any of its child structs or own struct???

func main() {

	// Create a Gin router with default middleware (logger and recovery)
	engine := gin.Default()

	fmt.Println("engine")
	repo := customer.NewMysqlUserRepository(db)
	fmt.Println("check repo")
	service := customer.NewUserService1(repo)
	fmt.Println("check service")
	handler := customer.NewUserHandler(service)
	fmt.Println("check handler")

	//Load HTML files and css
	engine.LoadHTMLGlob("web/html/*")
	engine.Static("/styles", "./web/styles")
	engine.Static("/icons", "./web/icons")

	tmpDbConfig()

	api.RegisterWebRouts(engine)
	api.RegisterApiRouts(engine, handler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}

func tmpDbConfig() {
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
