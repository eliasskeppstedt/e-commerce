package main

import (
	"database/sql"
	"ecommerce/duckyarmy/configs"
	"ecommerce/duckyarmy/internal/handler"
	"ecommerce/duckyarmy/internal/repository"
	"ecommerce/duckyarmy/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appCfg configs.Config
var db *sql.DB // where to store, in app config, any of its child structs or own struct???

func main() {

	// Create a Gin router with default middleware (logger and recovery)
	engine := gin.Default()

	//Load HTML files and css
	engine.LoadHTMLGlob("web/html/*")
	engine.Static("/styles", "./web/styles")

	engine.GET("/ping", func(ctx *gin.Context) {
		// Return JSON response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		fmt.Println("message: good")
	})

	//standard sidor
	engine.GET("/", func(ctx *gin.Context) {
		// Return HTTP response
		ctx.HTML(http.StatusOK, "homePage.html", gin.H{})
		fmt.Println("Hompage works")
	})

	engine.GET("/products", func(ctx *gin.Context) {
		// Return HTTP response
		ctx.HTML(http.StatusOK, "productsPage.html", gin.H{})
		fmt.Println("productspage works")
	})

	engine.GET("/cart", func(ctx *gin.Context) {
		// Return HTTP response
		ctx.HTML(http.StatusOK, "cartPage.html", gin.H{})
		fmt.Println("Hompage works")
	})

	engine.GET("/login", func(ctx *gin.Context) {
		// Return HTTP response
		ctx.HTML(http.StatusOK, "loginPage.html", gin.H{})
		fmt.Println("productspage works")
	})
	engine.GET("/register", func(ctx *gin.Context) {
		// Return HTTP response
		ctx.HTML(http.StatusOK, "registerPage.html", gin.H{})
		fmt.Println("productspage works")
	})

	//Register handling
	//Här registrerar sig användaren
	engine.POST("/register", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		fmt.Println("Username:", username)
		fmt.Println("Password:", password)

		ctx.HTML(http.StatusOK, "registerPage.html", gin.H{})
		fmt.Println("registerpage Post Working")
	})

	var err error
	appCfg, err = configs.LoadDbCfg()
	if err != nil {
		log.Fatal(err)
	}

	dsnCfg := configs.LoadDsnCfg(&appCfg.DB) // unnecesary functions in config?

	// Get a database handle.

	db, err = sql.Open("mysql", dsnCfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Printf("\n -- Connected!\n\n")

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("engine")
	repo := repository.NewUserRepository(db)
	fmt.Println("check repo")
	service := service.NewUserService(repo)
	fmt.Println("check service")
	handler := handler.NewUserHandler(service)
	fmt.Println("check handler")

	engine.GET("/user/:userID", handler.GetUsers)

}
