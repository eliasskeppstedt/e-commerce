package cart

import (
	"ecommerce/duckyarmy/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandler struct {
	service *cartService1
}

func NewCartHandler(s *cartService1) *CartHandler {
	return &CartHandler{s}
}

func (h *CartHandler) AddToCart(ctx *gin.Context) {
	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	var req struct {
		ProductID int `json:"product_id"`
	}

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	err = h.service.AddItem(userID, req.ProductID, 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to add product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product added to cart",
	})
}

func (h *CartHandler) GetCartProducts(ctx *gin.Context) {
	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		return
	}
	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	products, err := h.service.GetCartProducts(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *CartHandler) RemoveItem(ctx *gin.Context) {

	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		return
	}
	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	err = h.service.RemoveItem(userID, productID, 1)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "item removed"})
}
