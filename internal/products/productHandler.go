package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service productService
}

func NewProductHandler(s productService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetProducts(ctx *gin.Context) {
	productid := ctx.Param("productID")
	numid, err1 := strconv.Atoi(productid)
	if err1 != nil {
		fmt.Println("Felaktigt productID")
		return
	}
	product, err2 := h.service.getProductsByProductID(numid)
	if err2 != nil {
		fmt.Println("Någonting har gått fel i products_handler")
		return
	}
	ctx.JSON(http.StatusOK, product)
}
