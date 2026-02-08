package main

import (
	"database/sql"
	"ecommerce/duckyarmy/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	engine := gin.Default()

	engine.GET("/ping", func(ctx *gin.Context) {
		// Return JSON response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		fmt.Println("message: good")
	})

	var err error
	appCfg, err := config.LoadDbCfg()
	if err != nil {
		log.Fatal(err)
	}

	dsnCfg := config.LoadDsnCfg(&appCfg.DB) // unnecesary functions in config?

	// Get a database handle.
	var db *sql.DB // where to store, in app config, any of its child structs or own struct???
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
}
