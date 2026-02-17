package customer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service userService
}

func NewUserHandler(s userService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	userid := ctx.Param("userID")
	numid, err1 := strconv.Atoi(userid)
	if err1 != nil {
		fmt.Println("Felaktigt userID")
		return
	}
	user, err2 := h.service.getUsersByUserID(numid)
	if err2 != nil {
		fmt.Println("Någonting har gått fel i users_handler")
		return
	}
	ctx.JSON(http.StatusOK, user)
}
