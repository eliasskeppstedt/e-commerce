package main

import (
	"database/sql"
	"ecommerce/duckyarmy/api"
	"ecommerce/duckyarmy/internal/cart"
	"ecommerce/duckyarmy/internal/category"
	"ecommerce/duckyarmy/internal/customer"
	"ecommerce/duckyarmy/internal/order"
	"ecommerce/duckyarmy/internal/product"
	"ecommerce/duckyarmy/internal/transaction"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

func main() {

	// Create Gin router
	engine := gin.Default()
	engine.SetTrustedProxies(nil)

	// Initialize database
	db := tmpDbConfig()
	txManager := transaction.NewTxManager(db)

	goose.SetDialect("mysql")
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	// USER SETUP
	userRepo := customer.NewMysqlUserRepository(db)
	userService := customer.NewUserService1(userRepo)
	userHandler := customer.NewUserHandler(userService)

	// PRODUCT SETUP
	productRepo := product.NewMysqlProductRepository(db)
	productService := product.NewProductServiceImp(*txManager, productRepo)
	productHandler := product.NewProductHandler(productService)

	cartRepo := cart.NewMysqlCartRepository(db)
	cartService := cart.NewCartService1(*txManager, productRepo, cartRepo)
	cartHandler := cart.NewCartHandler(cartService)

	orderRepo := order.NewMysqlOrderRepository(db)
	orderService := order.NewOrderService1(*txManager, orderRepo, cartRepo, productRepo)
	orderHandler := order.NewOrderHandler(orderService)
	// CATEGORY SETUP
	categoryRepo := category.NewMysqlCategoryRepository(db)
	categoryService := category.NewCategoryServiceImp(categoryRepo)
	categoryHandler := category.NewCategoryHandler(categoryService)

	// Load HTML and static files
	engine.LoadHTMLGlob("web/html/*")
	engine.Static("/styles", "./web/styles")
	engine.Static("/icons", "./web/icons")
	engine.Static("/scripts", "./web/scripts")

	// Register routes
	api.RegisterWebRouts(engine)
	api.RegisterApiRouts(engine, userHandler, productHandler, cartHandler, orderHandler, categoryHandler)

	// Start server
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func tmpDbConfig() *sql.DB {

	// environment-variabler laddas in via docker-compose

	// MySQL config
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBURL")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.ParseTime = true

	// Open connection, säkerställ att dbn hinner starta
	for i := 0; i < 10; i++ {
		db, err := sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Println("Error opening database:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if err = db.Ping(); err == nil {
			fmt.Println("Connected to database")
			return db
		}

		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Could not connect to database after retries")
	return nil
}
