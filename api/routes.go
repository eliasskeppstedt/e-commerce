package api

import (
	"ecommerce/duckyarmy/internal/customer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWebRouts(engine *gin.Engine) {
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
}

func RegisterApiRouts(engine *gin.Engine, handler *customer.UserHandler) {
	// registrera handlers för kund
	fmt.Println("registering user handler")
	engine.POST("/api/register", handler.RegisterUser)

	engine.GET("/api/user/:userID", handler.GetUserByUsername)
}
