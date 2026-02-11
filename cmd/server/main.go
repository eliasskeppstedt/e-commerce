package main

import (
	"database/sql"
	"ecommerce/duckyarmy/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var appCfg configs.Config
var db *sql.DB // where to store, in app config, any of its child structs or own struct???

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

	test()
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}

type Users struct {
	userID       int
	userName     string
	password     string
	emailAddress string
	firstName    string
	lastName     string
	address      string
	zipCode      string
	phoneNumber  string
}

// albumByID queries for the album with the specified ID.
func test() {
	// An album to hold data from the returned row.
	var user Users
	id := 1
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&user.userID, &user.password, &user.emailAddress, &user.firstName, &user.lastName, &user.address, &user.zipCode, &user.phoneNumber); err != nil {
		if err == sql.ErrNoRows {
			//fmt.Errorf("albumsById %d: no such album", id)
			fmt.Println("error 1")
		}
		//fmt.Errorf("albumsById %d: %v", id, err)
		fmt.Println("error 2")
	}
	fmt.Println(row)
}
