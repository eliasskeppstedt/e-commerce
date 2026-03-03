package api

import (
	"ecommerce/duckyarmy/internal/auth"
	"ecommerce/duckyarmy/internal/customer"
	"ecommerce/duckyarmy/internal/product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRouts(engine *gin.Engine) {
	//standard sidor
	engine.GET("/", auth.Middleware(), func(ctx *gin.Context) {

		claimsValue, exists := ctx.Get("auth_token")
		if !exists {
			ctx.HTML(http.StatusOK, "homePage.html", gin.H{})
			fmt.Println("Homepage works")
			return
		}

		claims := claimsValue.(*auth.Claims)

		ctx.HTML(http.StatusOK, "homePage.html", gin.H{
			"UserID":  claims.UserID,
			"IsAdmin": claims.IsAdmin,
		})
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

func RegisterApiRouts(
	engine *gin.Engine,
	userHandler *customer.UserHandler,
	productHandler *product.ProductHandler,
) {

	// registrera handlers för kund
	fmt.Println("registering user handler")
	engine.POST("/api/users/register", userHandler.RegisterUser)
	engine.POST("/api/users/login", userHandler.UserLogin)
	engine.GET("/api/products", productHandler.GetProducts)
	//engine.POST("/api/products", productHandler.CreateProduct)
}
