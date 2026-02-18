package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service userService
}

func NewUserHandler(s userService1) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err2 := h.service.getUserByUsername(username)
	if err2 != nil {
		fmt.Println("Någonting har gått fel i users_handler")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) Register(ctx *gin.Context) {
	//implement
}
