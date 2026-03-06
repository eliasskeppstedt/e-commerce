package customer

import (
	"ecommerce/duckyarmy/internal/auth"
	"ecommerce/duckyarmy/internal/cart"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service  *userService1
	cService cart.CartService
}

func NewUserHandler(s *userService1, c cart.CartService) *UserHandler {
	return &UserHandler{service: s, cService: c}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	first_name := ctx.PostForm("first_name")
	last_name := ctx.PostForm("last_name")
	address := ctx.PostForm("address")
	zip_code := ctx.PostForm("zip_code")
	phone_number := ctx.PostForm("phone_number")

	fmt.Println("Login1 username, password =", username, password)
	userID, err := h.service.registerUser(username,
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
		return
	}

	err = h.cService.CreateCart(userID)
	fmt.Println("userRegister complete, error? : ", err)
	ctx.Redirect(http.StatusSeeOther, "/login")
}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	loginInput := ctx.PostForm("input_login")
	password := ctx.PostForm("password")

	userID, isAdmin, err := h.service.userLogin(loginInput, password)
	if err != nil {
		//Felaktig inloggning
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else if userID == -1 {
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else {
		//FIX DOMAIN to os.GETenv("SERVERURL")
		token, err1 := auth.GenerateToken(userID, isAdmin)
		fmt.Println("Token:", token)
		if err1 != nil {
			ctx.Set("auth_token", "")
			ctx.Redirect(http.StatusSeeOther, "/login")
			fmt.Println("error i 4")
		}
		ctx.SetCookie("auth_token", token, 3600, "/", "", false, true)
		ctx.Redirect(http.StatusSeeOther, "/")
	}

}
func (h *UserHandler) UserLogout(ctx *gin.Context) {
	ctx.SetCookie("auth_token", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/")
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	user, err := h.service.getUserByID(userID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user_id":    user.UserID,
		"username":   user.UserName,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"address":    user.Address,
		"zip_code":   user.ZipCode,
		"phone":      user.PhoneNumber,
	})

}
