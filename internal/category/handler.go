package category

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *categoryServiceImp
}

func NewCategoryHandler(s *categoryServiceImp) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) GetCategories(ctx *gin.Context) {
	categories, err := h.service.getAll()
	if err != nil {
		fmt.Println("GET ALL CATEGORIES ERROR:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	fmt.Println("CreateCategory called")
	var c Category

	if err := ctx.ShouldBindJSON(&c); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("CATEGORY RECEIVED: %+v\n", c)

	if err := h.service.createCategory(c); err != nil {
		fmt.Println("Error inserting category:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create category"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Category created"})

}

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.service.deleteCategory(id); err != nil {
		// Check for FK constraint
		if strings.Contains(err.Error(), "foreign key constraint") {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Cannot delete category with products"})
			return
		}
		fmt.Println("DELETE CATEGORY ERROR:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
