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
	//var user user

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	fmt.Println("username, password, email:", username, password, email)

	err := h.service.registerUser(username, password, email)
	if err != nil {
		fmt.Println("Problem med att registrera användare")
		ctx.HTML(http.StatusBadRequest, "registerPage.html", gin.H{})
	} else {
		fmt.Println("Registrera användare funkade")
		ctx.HTML(http.StatusOK, "loginPage.html", gin.H{})

		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		email := ctx.PostForm("email")
		fmt.Println("username, password, email:", username, password, email)

		err := h.service.registerUser(username, password, email)
		if err != nil {
			fmt.Println("Problem med att registrera användare")
			ctx.Redirect(http.StatusSeeOther, "/register")
		} else {

			fmt.Println("Problem med att registrera användare")
			ctx.Redirect(http.StatusSeeOther, "/register")

			fmt.Println("Register Success")
			ctx.Redirect(http.StatusSeeOther, "/login")

		}
	}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	fmt.Println("Login1 username, password =", username, password)
	err := h.service.registerUser(username, password, email)

	if err != nil {
		fmt.Println("ERROR UserLogin in handler:", err)
		ctx.HTML(http.StatusBadRequest, "loginPage.html", gin.H{})
	} else {
		fmt.Println("userLogin complete")
		ctx.HTML(http.StatusOK, "homePage.html", gin.H{})
	}

}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	loginInput := ctx.PostForm("input_login")
	password := ctx.PostForm("password")

	token, err := h.service.userLogin(loginInput, password)
	if err != nil {
		ctx.SetCookie("auth_token", "", 0, "/", "127.0.0.1", false, true)
	} else {
		//FIX DOMAIN to os.GETenv("SERVERURL")
		ctx.SetCookie("auth_token", token, 3600, "/", "127.0.0.1", false, true)
		ctx.Redirect(http.StatusSeeOther, "/")
	}

}
