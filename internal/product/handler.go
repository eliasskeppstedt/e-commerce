package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *productServiceImp
}

func NewProductHandler(s *productServiceImp) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := h.service.getAll()
	if err != nil {
		fmt.Println("GET ALL PRODUCTS ERROR:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	fmt.Println("CreateProduct called")
	var p Product

	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("PRODUCT RECEIVED: %+v\n", p)

	if err := h.service.registerProduct(p); err != nil {
		fmt.Println("Error inserting product:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) error {
	return nil
}
