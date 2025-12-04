package middleware

import (
	"fmt"
	"go-cinema-api/models/web"
	"go-cinema-api/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		// 1. Ambil Header Authorization
		authHeader := ctx.GetHeader("Authorization")
		
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Success: false,
				Message: "Authorization Header not found",
				Data: nil,
			})
			return
		}

		// 2. Cek format Authorization Header, apakah dimulai dengan "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Success: false,
				Message: "Invalid Authorization format, expected 'Bearer <token>'",
				Data: nil,
			})
			return
		}

		// 3. Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Parse & Validasi token
		token, err := jwt.ParseWithClaims(tokenString, &utils.JWTCustomClaims{}, func(token *jwt.Token) (interface {}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Success: false,
				Message: "Invalid token or expired token",
				Data: nil,
			})
			return
		}

		// 5. Ambil klaim dari token dan simpan ke context
		if claims, ok := token.Claims.(*utils.JWTCustomClaims); ok {
			ctx.Set("userID", claims.Subject)
			ctx.Set("userRole", claims.Role)
		}
		ctx.Next()

	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		// 1. Ambil userRole dari context
		userRole, exists := ctx.Get("userRole")

		// 2. Cek apakah userRole ada dan bernilai "admin"
		if !exists || userRole != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, web.WebResponse{
				Success: false,
				Message: "Access forbidden: Admins only",
				Data: nil,
			})
			return
		}
		
		ctx.Next()
	}
}