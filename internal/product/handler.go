package product

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	fmt.Println(id)
	if err := h.service.deleteProduct(id); err != nil {
		fmt.Println("DELETE PRODUCT ERROR:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData struct {
		Stock int     `json:"stock"`
		Price float64 `json:"price"`
	}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	err = h.service.updateProduct(id, updateData.Stock, updateData.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}
