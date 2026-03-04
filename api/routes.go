package api

import (
	"ecommerce/duckyarmy/internal/cart"
	"ecommerce/duckyarmy/internal/category"
	"ecommerce/duckyarmy/internal/customer"
	"ecommerce/duckyarmy/internal/order"
	"ecommerce/duckyarmy/internal/product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	engine.GET("/categories", func(ctx *gin.Context) {
		// Return HTTP response
		ctx.HTML(http.StatusOK, "categoriesPage.html", gin.H{})
		fmt.Println("categoriespage works")
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
	cartHandler *cart.CartHandler,
	orderHandler *order.OrderHandler,
	categoryHandler *category.CategoryHandler,

) {

	// registrera handlers för kund
	fmt.Println("registering user handler")
	engine.POST("/api/users/register", userHandler.CreateAccount)
	engine.GET("/api/users/:user_id", userHandler.GetUserByUsername)

	// handlers för produkter
	engine.GET("/api/products", productHandler.GetProducts)
	engine.POST("/api/products", productHandler.CreateProduct)

	engine.POST("/api/cart/items", cartHandler.AddItem)
	engine.POST("/api/cart/checkout", orderHandler.CheckOut)
	engine.DELETE("/api/products/:id", productHandler.DeleteProduct)
	engine.PUT("/api/products/:id", productHandler.UpdateProduct)

	// handlers för kategorier
	engine.GET("/api/categories", categoryHandler.GetCategories)
	engine.POST("/api/categories", categoryHandler.CreateCategory)
	engine.DELETE("/api/categories/:id", categoryHandler.DeleteCategory)

}
