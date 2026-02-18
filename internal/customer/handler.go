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

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	fmt.Println("in user handler register user")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	err := h.service.register(username, password)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.HTML(http.StatusOK, "registerPage.html", gin.H{})
	fmt.Println("registerpage Post Working")
}
