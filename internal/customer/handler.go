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
	user, err2 := h.service.getByID(numid)
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
