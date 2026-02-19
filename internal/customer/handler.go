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
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	fmt.Println("username, password, email:", username, password, email)

	err := h.service.registerUser(username, password, email)
	if err != nil {
		fmt.Println("Problem med att registrera användare")
		ctx.Redirect(http.StatusSeeOther, "/register")
	} else {
		fmt.Println("Register Success")
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
}
