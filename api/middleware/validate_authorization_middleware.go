package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"forum-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"os"
)

func ValidateAuthorizationMiddleware(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"error": gin.H{
				"message": "Authorization is required",
			},
		})
		ctx.Abort()
		return
	}

	token, err := utils.ValidateToken(authorization)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": gin.H{
				"message": "Invalid token",
			},
		})
		ctx.Abort()
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	ctx.Set("user_id", claims["userID"])
	ctx.Next()
}

func GenerateHmac(uuid string, date string) string {
	mac := hmac.New(sha256.New, []byte(os.Getenv("SECRET_TOKEN")))
	io.WriteString(mac, uuid)
	io.WriteString(mac, date)
	return hex.EncodeToString(mac.Sum(nil))
}
