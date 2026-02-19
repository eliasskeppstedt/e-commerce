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
	idStr := ctx.PostForm("product_id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Product_id",
		})
	}

	product, err2 := h.service.getByProductID(id)
	if err2 != nil {
		fmt.Println("Någonting har gått fel i products_handler")
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (s *ProductHandler) DeleteProduct(ctx *gin.Context) error {
	return nil
}
