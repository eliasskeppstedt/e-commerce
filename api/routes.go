package api

import (
	"ecommerce/duckyarmy/internal/auth"
	"ecommerce/duckyarmy/internal/cart"
	"ecommerce/duckyarmy/internal/category"
	"ecommerce/duckyarmy/internal/customer"
	"ecommerce/duckyarmy/internal/order"
	"ecommerce/duckyarmy/internal/product"
	"ecommerce/duckyarmy/internal/review"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRouts(engine *gin.Engine) {
	//standard sidor
	engine.GET("/", auth.Middleware(), func(ctx *gin.Context) {
		claimsValue, exists := ctx.Get("auth_token")
		if exists {
			claims := claimsValue.(*auth.Claims)
			fmt.Println("claims.UserID", claims.UserID)
			ctx.HTML(http.StatusOK, "homePage.html", gin.H{
				"UserID":  claims.UserID,
				"IsAdmin": claims.IsAdmin})
			return
		}
		ctx.HTML(http.StatusOK, "homePage.html", gin.H{})
		fmt.Println("Homepage works")
	})

	engine.GET("/products", auth.Middleware(), func(ctx *gin.Context) {
		claimsValue, exists := ctx.Get("auth_token")
		if exists {
			claims := claimsValue.(*auth.Claims)
			fmt.Println("claims.UserID", claims.UserID)
			ctx.HTML(http.StatusOK, "productsPage.html", gin.H{
				"UserID":  claims.UserID,
				"IsAdmin": claims.IsAdmin})
			return
		}
		ctx.HTML(http.StatusOK, "productsPage.html", gin.H{})
		fmt.Println("Productpage works")
	})

	engine.GET("/cart", auth.Middleware(), func(ctx *gin.Context) {
		claimsValue, exists := ctx.Get("auth_token")
		if exists {
			claims := claimsValue.(*auth.Claims)
			fmt.Println("claims.UserID", claims.UserID)
			ctx.HTML(http.StatusOK, "cartPage.html", gin.H{
				"UserID":  claims.UserID,
				"IsAdmin": claims.IsAdmin})
			return
		}
		ctx.HTML(http.StatusUnauthorized, "loginPage.html", gin.H{})
		fmt.Println("Productpage works")
	})

	engine.GET("/profile", auth.Middleware(), func(ctx *gin.Context) {
		claimsValue, exists := ctx.Get("auth_token")
		if exists {
			claims := claimsValue.(*auth.Claims)
			fmt.Println("claims.UserID", claims.UserID)
			ctx.HTML(http.StatusOK, "profilePage.html", gin.H{
				"UserID":  claims.UserID,
				"IsAdmin": claims.IsAdmin})
			return
		}
		ctx.HTML(http.StatusUnauthorized, "loginPage.html", gin.H{})
		fmt.Println("Productpage works")
	})

	engine.GET("/order", auth.Middleware(), func(ctx *gin.Context) {
		claimsValue, exists := ctx.Get("auth_token")
		if exists {
			claims := claimsValue.(*auth.Claims)
			fmt.Println("claims.UserID", claims.UserID)
			ctx.HTML(http.StatusOK, "orderPage.html", gin.H{
				"UserID":  claims.UserID,
				"IsAdmin": claims.IsAdmin})
			return
		}
		ctx.HTML(http.StatusUnauthorized, "loginPage.html", gin.H{})
		fmt.Println("Productpage works")
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

	engine.GET("/orders", auth.Middleware(), func(ctx *gin.Context) {

		claimsValue, _ := ctx.Get("auth_token")
		claims := claimsValue.(*auth.Claims)

		ctx.HTML(http.StatusOK, "ordersPage.html", gin.H{
			"UserID": claims.UserID,
		})
	})
}

func RegisterApiRouts(
	engine *gin.Engine,
	userHandler *customer.UserHandler,
	productHandler *product.ProductHandler,
	cartHandler *cart.CartHandler,
	orderHandler *order.OrderHandler,
	categoryHandler *category.CategoryHandler,
	reviewHandler *review.ReviewHandler,

) {

	// registrera handlers för kund
	fmt.Println("registering user handler")
	engine.POST("/api/users/register", userHandler.RegisterUser)
	engine.POST("/api/users/login", userHandler.UserLogin)
	engine.GET("/api/users/logout", userHandler.UserLogout)
	engine.GET("/api/users/profile", auth.Middleware(), userHandler.GetUserByID)

	// handlers för produkter
	engine.GET("/api/products", productHandler.GetProducts)
	engine.POST("/api/products", productHandler.CreateProduct)
	engine.DELETE("/api/products/:id", productHandler.DeleteProduct)
	engine.PUT("/api/products/:id", productHandler.UpdateProduct)

	engine.GET("/api/cart", auth.Middleware(), cartHandler.GetCartProducts)
	engine.POST("api/product/add", auth.Middleware(), cartHandler.AddToCart)

	engine.DELETE("/api/cart/items/:product_id", auth.Middleware(), cartHandler.RemoveItem)

	// handlers för kategorier
	engine.GET("/api/categories", categoryHandler.GetCategories)
	engine.POST("/api/categories", categoryHandler.CreateCategory)
	engine.DELETE("/api/categories/:id", categoryHandler.DeleteCategory)

	engine.POST("/api/cart/checkout", auth.Middleware(), orderHandler.CheckOut)

	engine.GET("/api/reviews/:product_id", reviewHandler.GetReviews)
	engine.POST("/api/reviews", auth.Middleware(), reviewHandler.AddReview)
	engine.DELETE("/api/reviews/:id", auth.Middleware(), reviewHandler.DeleteReview)

	engine.GET("/api/orders", auth.Middleware(), orderHandler.GetOrders)

	//engine.POST("/api/products", productHandler.CreateProduct)
}
