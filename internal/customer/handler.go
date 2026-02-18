package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *userService1
}

func NewUserHandler(s *userService1) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	user, err := h.service.getUserByUsername(username)
	if err != nil {
		fmt.Println("Någonting har gått fel i users_handler")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateAccount(ctx *gin.Context) {
	username := ctx.Param("username")
	password := ctx.Param("password")
	emailaddress := ctx.Param("emailaddress")

	err := h.service.registerUser(username, password, emailaddress)
	if err != nil {
		fmt.Println("Problem med att registrera användare")
	}
}
