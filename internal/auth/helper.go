package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// returnerar -1 om userID inte hittas, sätter json error response
func GetUserID(ctx *gin.Context) int {
	claimsValue, exists := ctx.Get("auth_token")

	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "user not authenticated"})
		return -1
	}

	claims := claimsValue.(*Claims)
	return claims.UserID
}
