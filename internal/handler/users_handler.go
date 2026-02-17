package handler

import (
	"ecommerce/duckyarmy/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UsersService
}

func NewUserHandler(s service.UsersService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	userid := ctx.Param("userID")
	numid, err1 := strconv.Atoi(userid)
	if err1 != nil {
		fmt.Println("Felaktigt userID")
		return
	}
	user, err2 := h.service.GetUsersByUserID(numid)
	if err2 != nil {
		fmt.Println("Någonting har gått fel i users_handler")
		return
	}
	ctx.JSON(http.StatusOK, user)
}
