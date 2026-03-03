package customer

import (
	"ecommerce/duckyarmy/internal/auth"
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

func (h *UserHandler) RegisterUser(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	first_name := ctx.PostForm("first_name")
	last_name := ctx.PostForm("lastname_name")
	address := ctx.PostForm("address")
	zip_code := ctx.PostForm("zip_code")
	phone_number := ctx.PostForm("phone_number")

	fmt.Println("Login1 username, password =", username, password)
	err := h.service.registerUser(username,
		password,
		email,
		first_name,
		last_name,
		address,
		zip_code,
		phone_number)

	if err != nil {
		fmt.Println("ERROR UserLogin in handler:", err)
		ctx.Redirect(http.StatusSeeOther, "/register")
	} else {
		fmt.Println("userLogin complete")
		ctx.Redirect(http.StatusSeeOther, "/login")
	}

}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	loginInput := ctx.PostForm("input_login")
	password := ctx.PostForm("password")

	userID, isAdmin, err := h.service.userLogin(loginInput, password)
	if err != nil {
		ctx.Set("auth_token", "")
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else if userID == -1 {
		ctx.Set("auth_token", "")
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else {
		//FIX DOMAIN to os.GETenv("SERVERURL")
		token, err1 := auth.GenerateToken(userID, isAdmin)
		if err1 != nil {
			ctx.Set("auth_token", "")
			ctx.Redirect(http.StatusSeeOther, "/login")
		}
		ctx.SetCookie("auth_token", token, 3600, "/", "127.0.0.1", false, true)
		ctx.Redirect(http.StatusSeeOther, "/")
	}

}
