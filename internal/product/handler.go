package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service productService
}

func NewProductHandler(s productService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetProducts(ctx *gin.Context) {
	productId := ctx.PostForm("product_id")

	product, err2 := h.service.getByProductID(productId)
	if err2 != nil {
		fmt.Println("Någonting har gått fel i products_handler")
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (s *ProductHandler) DeleteProduct(ctx *gin.Context) error {
	return nil
}
